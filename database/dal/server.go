package dal

import (
	"database/sql"
	"time"

	verrors "github.com/akrck02/whisper/sdk/errors"
	"github.com/akrck02/whisper/sdk/models"
)

func CreateServer(db *sql.DB, server *models.Server) (*int64, *verrors.VError) {
	if nil == db {
		return nil, verrors.Unexpected(verrors.DatabaseConnectionEmptyMessage)
	}

	if nil == server {
		return nil, verrors.InvalidRequest(verrors.ServerOwnerEmptyMessage)
	}

	// TODO: Check that server owner exists

	statement, err := db.Prepare(
		"INSERT INTO server(owner, name, description, profile_pic, insert_date) VALUES(?,?,?,?,?)",
	)

	if nil != err {
		return nil, verrors.DatabaseError(err.Error())
	}

	res, err := statement.Exec(
		server.Owner,
		server.Name,
		server.ProfilePic,
		server.Description,
		time.Now(),
	)

	if nil != err {
		return nil, verrors.DatabaseError(err.Error())
	}

	server.ID, err = res.LastInsertId()
	if nil != err {
		return nil, verrors.DatabaseError(err.Error())
	}

	return &server.ID, nil

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
		WHERE id=?
	`)

	if nil != err {
		return verrors.DatabaseError(err.Error())
	}

	res, err := statement.Exec(
		server.Name,
		server.Description,
		server.ProfilePic,
		server.ID,
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

func GetServer(db *sql.DB, id int64) (*models.Server, *verrors.VError) {

	if nil == db {
		return nil, verrors.Unexpected(verrors.DatabaseConnectionEmptyMessage)
	}

	statement, err := db.Prepare(`
		SELECT user_agent,
			owner_id,
			name,
			description,
			profile_pic,
			insert_date
		FROM server
		WHERE user_id = ?
	`)
	if nil != err {
		return nil, verrors.DatabaseError(err.Error())
	}

	rows, err := statement.Query(id)
	if nil != err {
		return nil, verrors.DatabaseError(err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, verrors.NotFound("Server not found.")
	}

	var ownerId int64
	var name string
	var description string
	var profilePic string
	var insertDate int64
	rows.Scan(&id, &ownerId, &name, &description, &profilePic, &insertDate)

	return &models.Server{
		ID:          id,
		Owner:       ownerId,
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
