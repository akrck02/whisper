package dal

import (
	"database/sql"
	"time"

	"github.com/akrck02/whisper/sdk/cryptography"
	verrors "github.com/akrck02/whisper/sdk/errors"
	"github.com/akrck02/whisper/sdk/models"
	"github.com/akrck02/whisper/sdk/validations"
	"github.com/google/uuid"
)

func CreateUser(db *sql.DB, user *models.User) (*int64, *verrors.VError) {

	if nil == db {
		return nil, verrors.Unexpected(verrors.DatabaseConnectionEmptyMessage)
	}

	if nil == user {
		return nil, verrors.InvalidRequest(verrors.UserNotFoundMessage)
	}

	err := validations.ValidateEmail(user.Email)
	if nil != err {
		return nil, verrors.InvalidRequest(err.Error())
	}

	if user.Username == "" {
		return nil, verrors.InvalidRequest(verrors.UserNameEmptyMessage)
	}

	err = validations.ValidatePassword(user.Password)
	if nil != err {
		return nil, verrors.InvalidRequest(err.Error())
	}

	hashedPassword, err := cryptography.Hash(user.Password)
	if nil != err {
		return nil, verrors.Unexpected(err.Error())
	}

	// Check email
	usr, userGetErr := GetUserByEmail(db, user.Email)
	if nil == userGetErr && nil != usr {
		return nil, verrors.New(verrors.UserAlreadyExistsErrorCode, verrors.UserAlreadyExistsMessage)
	}

	// Check username
	usr, userGetErr = GetUserByUsername(db, user.Username)
	if nil == userGetErr && nil != usr {
		return nil, verrors.New(verrors.UserAlreadyExistsErrorCode, verrors.UserAlreadyExistsMessage)
	}

	statement, err := db.Prepare(`
		INSERT INTO user(
			uuid,
			email,
			username,
			profile_pic,
			password,
			insert_date
		) VALUES(?,?,?,?,?,?)
	`)

	if nil != err {
		return nil, verrors.DatabaseError(err.Error())
	}

	res, err := statement.Exec(
		uuid.NewString(),
		user.Email,
		user.Username,
		user.ProfilePicture,
		hashedPassword,
		time.Now().UnixMilli(),
	)

	if nil != err {
		return nil, verrors.DatabaseError(err.Error())
	}

	user.ID, err = res.LastInsertId()
	if nil != err {
		return nil, verrors.DatabaseError(err.Error())
	}

	return &user.ID, nil
}

func UpdateUserEmail(db *sql.DB, userId int64, email string) *verrors.VError {

	if nil == db {
		return verrors.Unexpected(verrors.DatabaseConnectionEmptyMessage)
	}

	err := validations.ValidateIsPositive(userId, "user id")
	if err != nil {
		return verrors.InvalidRequest(err.Error())
	}

	err = validations.ValidateEmail(email)
	if err != nil {
		return verrors.InvalidRequest(err.Error())
	}

	usr, userGetErr := GetUserByEmail(db, email)
	if nil == userGetErr && nil != usr {
		return verrors.New(verrors.UserAlreadyExistsErrorCode, verrors.UserAlreadyExistsMessage)
	}

	statement, err := db.Prepare(`
		UPDATE user
		SET email=?
		WHERE id=?
	`)

	res, err := statement.Exec(email, userId)
	if nil != err {
		return verrors.DatabaseError(err.Error())
	}

	affectedRows, err := res.RowsAffected()
	if nil != err {
		return verrors.DatabaseError(err.Error())
	}

	if affectedRows == 0 {
		return verrors.DatabaseError(verrors.UserCannotUpdateMessage)
	}

	return nil
}

func UpdateUserUsername(db *sql.DB, userId int64, username string) *verrors.VError {

	if nil == db {
		return verrors.Unexpected(verrors.DatabaseConnectionEmptyMessage)
	}

	err := validations.ValidateIsPositive(userId, "user id")
	if err != nil {
		return verrors.InvalidRequest(err.Error())
	}

	usr, userGetErr := GetUserByUsername(db, username)
	if nil == userGetErr && nil != usr {
		return verrors.New(verrors.UserAlreadyExistsErrorCode, verrors.UserAlreadyExistsMessage)
	}

	statement, err := db.Prepare(`
		UPDATE user
		SET username=?
		WHERE id=?
	`)

	res, err := statement.Exec(username, userId)
	if nil != err {
		return verrors.DatabaseError(err.Error())
	}

	affectedRows, err := res.RowsAffected()
	if nil != err {
		return verrors.DatabaseError(err.Error())
	}

	if affectedRows == 0 {
		return verrors.DatabaseError(verrors.UserCannotUpdateMessage)
	}

	return nil
}

func UpdateUserPassword(db *sql.DB, userId int64, password string) *verrors.VError {

	if nil == db {
		return verrors.Unexpected(verrors.DatabaseConnectionEmptyMessage)
	}

	err := validations.ValidateIsPositive(userId, "user id")
	if err != nil {
		return verrors.InvalidRequest(err.Error())
	}

	err = validations.ValidatePassword(password)
	if nil != err {
		return verrors.InvalidRequest(err.Error())
	}

	hashedPassword, err := cryptography.Hash(password)
	if nil != err {
		return verrors.Unexpected(err.Error())
	}

	statement, err := db.Prepare(`
		UPDATE user
		SET username=?
		WHERE id=?
	`)

	res, err := statement.Exec(hashedPassword, userId)
	if nil != err {
		return verrors.DatabaseError(err.Error())
	}

	affectedRows, err := res.RowsAffected()
	if nil != err {
		return verrors.DatabaseError(err.Error())
	}

	if affectedRows == 0 {
		return verrors.DatabaseError(verrors.UserCannotUpdateMessage)
	}

	return nil
}

func GetUser(db *sql.DB, id int64) (*models.User, *verrors.VError) {

	if nil == db {
		return nil, verrors.Unexpected(verrors.DatabaseConnectionEmptyMessage)
	}

	err := validations.ValidateIsPositive(id, "user id")
	if nil != err {
		return nil, verrors.InvalidRequest(err.Error())
	}

	statement, err := db.Prepare(`
		SELECT
			uuid,
			email,
			username,
			profile_pic,
			password,
			insert_date
		FROM user
		WHERE id=?
	`)

	rows, err := statement.Query(id)
	if nil != err {
		return nil, verrors.DatabaseError(err.Error())
	}

	if !rows.Next() {
		return nil, verrors.NotFound(verrors.UserNotFoundMessage)
	}
	defer rows.Close()

	var uuid string
	var email string
	var username string
	var profilePic string
	var password string
	var insertDate int64

	rows.Scan(
		&uuid,
		&email,
		&username,
		&profilePic,
		&password,
		&insertDate,
	)

	return &models.User{
		ID:             id,
		UUID:           uuid,
		Email:          email,
		Username:       username,
		ProfilePicture: profilePic,
		Password:       password,
		InsertDate:     insertDate,
	}, nil
}

func GetUserByEmail(db *sql.DB, email string) (*models.User, *verrors.VError) {

	if nil == db {
		return nil, verrors.Unexpected(verrors.DatabaseConnectionEmptyMessage)
	}

	statement, err := db.Prepare(`
		SELECT
			id,
			uuid,
			username,
			profile_pic,
			password,
			insert_date
		FROM user
		WHERE email=?
	`)

	rows, err := statement.Query(email)
	if nil != err {
		return nil, verrors.DatabaseError(err.Error())
	}

	if !rows.Next() {
		return nil, verrors.NotFound(verrors.UserNotFoundMessage)
	}
	defer rows.Close()

	var id int64
	var uuid string
	var username string
	var profilePic string
	var password string
	var insertDate int64

	rows.Scan(
		&id,
		&uuid,
		&username,
		&profilePic,
		&password,
		&insertDate,
	)

	return &models.User{
		ID:             id,
		UUID:           uuid,
		Email:          email,
		Username:       username,
		ProfilePicture: profilePic,
		Password:       password,
		InsertDate:     insertDate,
	}, nil
}

func GetUserByUsername(db *sql.DB, username string) (*models.User, *verrors.VError) {

	if nil == db {
		return nil, verrors.Unexpected(verrors.DatabaseConnectionEmptyMessage)
	}

	statement, err := db.Prepare(`
		SELECT
			id,
			uuid,
			email,
			profile_pic,
			password,
			insert_date
		FROM user
		WHERE username=?
	`)

	rows, err := statement.Query(username)
	if nil != err {
		return nil, verrors.DatabaseError(err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, verrors.NotFound(verrors.UserNotFoundMessage)
	}

	var id int64
	var uuid string
	var email string
	var profilePic string
	var password string
	var insertDate int64

	rows.Scan(
		&id,
		&uuid,
		&email,
		&profilePic,
		&password,
		&insertDate,
	)

	return &models.User{
		ID:             id,
		UUID:           uuid,
		Email:          email,
		Username:       username,
		ProfilePicture: profilePic,
		Password:       password,
		InsertDate:     insertDate,
	}, nil
}

func DeleteUser(db *sql.DB, id int64) *verrors.VError {

	if nil == db {
		return verrors.Unexpected(verrors.DatabaseConnectionEmptyMessage)
	}

	statement, err := db.Prepare(`DELETE FROM user WHERE id=?`)
	res, err := statement.Exec(id)
	if nil != err {
		return verrors.DatabaseError(err.Error())
	}

	affectedRows, err := res.RowsAffected()
	if nil != err {
		return verrors.DatabaseError(err.Error())
	}

	if affectedRows == 0 {
		return verrors.DatabaseError(verrors.UserCannotDeleteMessage)
	}

	return nil
}
