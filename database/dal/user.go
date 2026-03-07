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
			password
		) VALUES(?,?,?,?,?)
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
		time.Now(),
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

func UpdateUser(db *sql.DB) *verrors.VError {
	err := verrors.TODO()
	return &err
}

func GetUser(db *sql.DB, id int64) (*models.User, *verrors.VError) {
	err := verrors.TODO()
	return nil, &err
}

func GetUserByEmail(db *sql.DB, email string) (*models.User, *verrors.VError) {
	err := verrors.TODO()
	return nil, &err
}

func GetUserByUsername(db *sql.DB, username string) (*models.User, *verrors.VError) {
	err := verrors.TODO()
	return nil, &err
}

func DeleteUser(db *sql.DB, id int64) *verrors.VError {
	err := verrors.TODO()
	return &err
}
