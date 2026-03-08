package dal

import (
	"database/sql"
	"time"

	verrors "github.com/akrck02/whisper/sdk/errors"
	"github.com/akrck02/whisper/sdk/models"
	"github.com/akrck02/whisper/sdk/validations"
	"github.com/google/uuid"
)

func CreateServer(mainDb *sql.DB, serverDb *sql.DB, ownerId int64, server *models.Server) (string, *verrors.VError) {

	if nil == mainDb {
		return "", verrors.Unexpected(verrors.DatabaseConnectionEmptyMessage)
	}

	if nil == serverDb {
		return "", verrors.Unexpected(verrors.DatabaseConnectionEmptyMessage)
	}

	if 0 == ownerId {
		return "", verrors.InvalidRequest(verrors.ServerOwnerEmptyMessage)
	}

	err := validations.ValidateIsPositive(ownerId, "server owner id")
	if nil != err {
		return "", verrors.InvalidRequest(err.Error())
	}

	usr, usrGetErr := GetUser(mainDb, ownerId)
	if nil != usrGetErr || nil == usr {
		return "", verrors.InvalidRequest(verrors.ServerOwnerNotFoundMessage)
	}

	if nil == server {
		return "", verrors.InvalidRequest(verrors.ServerEmptyMessage)
	}

	if "" == server.Name {
		return "", verrors.InvalidRequest(verrors.ServerNameEmptyMessage)
	}

	statement, err := serverDb.Prepare(
		"INSERT INTO server (uuid, name, description, profile_pic, insert_date) VALUES(?,?,?,?,?)",
	)

	if nil != err {
		return "", verrors.DatabaseError(err.Error())
	}

	newUuid := uuid.New()
	res, err := statement.Exec(
		newUuid,
		server.Name,
		server.Description,
		server.ProfilePic,
		time.Now().UnixMilli(),
	)

	if nil != err {
		return "", verrors.DatabaseError(err.Error())
	}

	affectedRows, err := res.RowsAffected()
	if nil != err {
		return "", verrors.DatabaseError(err.Error())
	}

	if affectedRows == 0 {
		return "", verrors.DatabaseError("")
	}

	server.UUID = newUuid.String()
	return server.UUID, nil

}

func UpdateServer(db *sql.DB, server *models.Server) *verrors.VError {

	if nil == db {
		return verrors.Unexpected(verrors.DatabaseConnectionEmptyMessage)
	}

	if nil == server {
		return verrors.InvalidRequest(verrors.ServerOwnerEmptyMessage)
	}

	statement, err := db.Prepare(`
		UPDATE server
		SET name=?,
		description=?,
		profile_pic=?
		WHERE uuid=?
	`)

	if nil != err {
		return verrors.DatabaseError(err.Error())
	}

	res, err := statement.Exec(
		server.Name,
		server.Description,
		server.ProfilePic,
		server.UUID,
	)
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

func GetServer(db *sql.DB, uuid string) (*models.Server, *verrors.VError) {

	if nil == db {
		return nil, verrors.Unexpected(verrors.DatabaseConnectionEmptyMessage)
	}

	statement, err := db.Prepare(`
		SELECT
			name,
			description,
			profile_pic,
			insert_date
		FROM server
		WHERE uuid = ?
	`)
	if nil != err {
		return nil, verrors.DatabaseError(err.Error())
	}

	rows, err := statement.Query(uuid)
	if nil != err {
		return nil, verrors.DatabaseError(err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, verrors.NotFound(verrors.ServerNotFoundError)
	}

	var name string
	var description string
	var profilePic string
	var insertDate int64
	rows.Scan(&name, &description, &profilePic, &insertDate)

	return &models.Server{
		UUID:        uuid,
		Name:        name,
		Description: description,
		ProfilePic:  profilePic,
		InsertDate:  insertDate,
	}, nil

}

func DeleteServer(db *sql.DB, id int64) *verrors.VError {

	if nil == db {
		return verrors.Unexpected(verrors.DatabaseConnectionEmptyMessage)
	}

	statement, err := db.Prepare(
		"DELETE FROM server WHERE id=?",
	)

	if nil != err {
		return verrors.DatabaseError(err.Error())
	}

	res, err := statement.Exec(id)

	affectedRows, err := res.RowsAffected()
	if nil != err {
		return verrors.DatabaseError(err.Error())
	}

	if affectedRows == 0 {
		return verrors.DatabaseError(verrors.UserCannotDeleteMessage)
	}

	return nil
}

func GetAllUserServers(db *sql.DB, userId int64) (*[]models.Server, *verrors.VError) {
	return nil, nil
}
