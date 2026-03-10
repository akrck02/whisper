package services

import (
	"database/sql"

	"github.com/akrck02/whisper/database/dal"
	verrors "github.com/akrck02/whisper/sdk/errors"
	"github.com/akrck02/whisper/sdk/models"
)

func CreateUser(db *sql.DB, user models.User) (string, *verrors.VError) {
	tx, err := db.Begin()
	if err != nil {
		return "", verrors.DatabaseError(err.Error())
	}

	defer tx.Rollback()

	userUuid, rerr := dal.CreateUser(db, &user)
	if nil != rerr {
		return "", rerr
	}

	return userUuid, nil
}

func GetUser(db *sql.DB, uuid string) (*models.User, *verrors.VError) {

	user, err := dal.GetUserByUuid(db, uuid)
	if nil != err {
		return nil, err
	}

	return user, nil
}

func GetUserByEmail(db *sql.DB, email string) (*models.User, *verrors.VError) {

	user, err := dal.GetUserByEmail(db, email)
	if nil != err {
		return nil, err
	}

	return user, nil
}
