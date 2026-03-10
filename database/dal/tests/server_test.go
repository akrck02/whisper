package tests

import (
	"database/sql"
	"testing"

	"github.com/akrck02/whisper/database/dal"
	"github.com/akrck02/whisper/database/tables"
	verrors "github.com/akrck02/whisper/sdk/errors"
	"github.com/akrck02/whisper/sdk/models"
)

func TestServerCrud(t *testing.T) {

	mainDb, err := NewTestDatabase(tables.MainDatabase)
	AssertVErrorDoesNotExist(t, err)
	defer mainDb.Close()

	serverDb, err := NewTestDatabase(tables.ServerDatabase)
	AssertVErrorDoesNotExist(t, err)
	defer serverDb.Close()

	userUuid, err := dal.CreateUser(mainDb, &models.User{
		Email:    "user@whisper.org",
		Username: "user01",
		Password: "$P4ssw0rdW3db1128",
	})
	AssertVErrorDoesNotExist(t, err)

	ownerUser, err := dal.GetUserByUuid(mainDb, userUuid)
	AssertVErrorDoesNotExist(t, err)

	expectedServer := &models.Server{
		Name:        "Ghosts and tales!",
		Description: "The scariest of servers.",
	}

	t.Run("Create server validations", func(t *testing.T) { createServerValidations(t, mainDb, serverDb, ownerUser) })
	t.Run("Create server", func(t *testing.T) { expectedServer = createServer(t, mainDb, serverDb, ownerUser, expectedServer) })

	t.Run("Get server validations", func(t *testing.T) { getServerValidations(t, serverDb) })
}

func createServerValidations(t *testing.T, mainDb *sql.DB, serverDb *sql.DB, ownerUser *models.User) {
	_, err := dal.CreateServer(nil, nil, 0, nil)
	AssertVError(t, err, verrors.UnexpectedErrorCode, verrors.DatabaseConnectionEmptyMessage)

	_, err = dal.CreateServer(mainDb, nil, 0, nil)
	AssertVError(t, err, verrors.UnexpectedErrorCode, verrors.DatabaseConnectionEmptyMessage)

	_, err = dal.CreateServer(mainDb, serverDb, 0, nil)
	AssertVError(t, err, verrors.InvalidRequestErrorCode, verrors.ServerOwnerEmptyMessage)

	_, err = dal.CreateServer(mainDb, serverDb, -1, nil)
	AssertVError(t, err, verrors.InvalidRequestErrorCode, verrors.ServerOwnerNegativeMessage)

	_, err = dal.CreateServer(mainDb, serverDb, 999, nil)
	AssertVError(t, err, verrors.InvalidRequestErrorCode, verrors.ServerOwnerNotFoundMessage)

	_, err = dal.CreateServer(mainDb, serverDb, ownerUser.ID, nil)
	AssertVError(t, err, verrors.InvalidRequestErrorCode, verrors.ServerEmptyMessage)

	_, err = dal.CreateServer(mainDb, serverDb, ownerUser.ID, &models.Server{})
	AssertVError(t, err, verrors.InvalidRequestErrorCode, verrors.ServerNameEmptyMessage)

}

func createServer(t *testing.T, mainDb *sql.DB, serverDb *sql.DB, owner *models.User, expectedServer *models.Server) *models.Server {

	serverUUID, err := dal.CreateServer(mainDb, serverDb, owner.ID, expectedServer)
	AssertVErrorDoesNotExist(t, err)

	obtainedServer, err := dal.GetServer(serverDb, serverUUID)
	AssertVErrorDoesNotExist(t, err)

	Assert(
		t,
		nil != obtainedServer &&
			expectedServer.Name == obtainedServer.Name &&
			expectedServer.Description == obtainedServer.Description,
		"expected and obtained server mismatch",
	)
	return obtainedServer

}

func getServerValidations(t *testing.T, db *sql.DB) {

}
