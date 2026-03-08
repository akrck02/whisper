package tests

import (
	"database/sql"
	"testing"

	"github.com/akrck02/whisper/database/dal"
	verrors "github.com/akrck02/whisper/sdk/errors"
	"github.com/akrck02/whisper/sdk/models"
)

func TestUserCrud(t *testing.T) {
	db, err := NewTestDatabase()
	AssertVErrorDoesNotExist(t, err)
	defer db.Close()

	expectedUser := &models.User{
		Email:    "user@whisper.org",
		Username: "user01",
		Password: "$P4ssw0rdW3db1128",
	}

	t.Run("Register validations", func(t *testing.T) { createUserValidations(t, db) })
	t.Run("Register user", func(t *testing.T) { expectedUser = createUser(t, db, expectedUser) })

	t.Run("Get user validations", func(t *testing.T) { getUserValidations(t, db) })

	t.Run("Update user mail validation errors", func(t *testing.T) { updateEmailValidations(t, db) })
	t.Run("Update user mail", func(t *testing.T) { expectedUser = updateUserEmail(t, db, expectedUser) })
}

func createUserValidations(t *testing.T, db *sql.DB) {
	_, err := dal.CreateUser(nil, nil)
	AssertVError(t, err, verrors.UnexpectedErrorCode, verrors.DatabaseConnectionEmptyMessage)

	_, err = dal.CreateUser(db, nil)
	AssertVError(t, err, verrors.InvalidRequestErrorCode, verrors.UserNotFoundMessage)

	_, err = dal.CreateUser(db, &models.User{})
	AssertVError(t, err, verrors.InvalidRequestErrorCode, verrors.EmailEmptyMessage)

	_, err = dal.CreateUser(db, &models.User{Email: "userwhisper.org"})
	AssertVError(t, err, verrors.InvalidRequestErrorCode, "mail: missing '@' or angle-addr")

	_, err = dal.CreateUser(db, &models.User{Email: "u@.org"})
	AssertVError(t, err, verrors.InvalidRequestErrorCode, "mail: missing '@' or angle-addr")

	_, err = dal.CreateUser(db, &models.User{Email: "user@whisper.org"})
	AssertVError(t, err, verrors.InvalidRequestErrorCode, verrors.UserNameEmptyMessage)

	_, err = dal.CreateUser(db, &models.User{Email: "user@whisper.org", Username: "usr01"})
	AssertVError(t, err, verrors.InvalidRequestErrorCode, verrors.PasswordEmptyMessage)

	_, err = dal.CreateUser(db, &models.User{Email: "user@whisper.org", Username: "usr01", Password: ""})
	AssertVError(t, err, verrors.InvalidRequestErrorCode, verrors.PasswordEmptyMessage)

	_, err = dal.CreateUser(db, &models.User{Email: "user@whisper.org", Username: "usr01", Password: "abcdefghijklmnñopq"})
	AssertVError(t, err, verrors.InvalidRequestErrorCode, verrors.PasswordNoNumericMessage)

	_, err = dal.CreateUser(db, &models.User{Email: "user@whisper.org", Username: "usr01", Password: "1bcdefghijklmnñopq"})
	AssertVError(t, err, verrors.InvalidRequestErrorCode, verrors.PasswordNoSpecialCharacterMessage)

	_, err = dal.CreateUser(db, &models.User{Email: "user@whisper.org", Username: "usr01", Password: "#1bcdefghijklmnñop"})
	AssertVError(t, err, verrors.InvalidRequestErrorCode, verrors.PasswordNoUppercaseCharacterMessage)

	_, err = dal.CreateUser(db, &models.User{Email: "user@whisper.org", Username: "usr01", Password: "#1BCDEFGHUJKLMNÑOP"})
	AssertVError(t, err, verrors.InvalidRequestErrorCode, verrors.PasswordNoLowercaseCharacterMessage)
}

func createUser(t *testing.T, db *sql.DB, user *models.User) *models.User {
	userID, err := dal.CreateUser(db, user)
	AssertVErrorDoesNotExist(t, err)

	obtainedUser, err := dal.GetUser(db, *userID)
	AssertVErrorDoesNotExist(t, err)
	Assert(t, nil != obtainedUser && user.Email == obtainedUser.Email, "expected user and obtained user mismatch")
	return obtainedUser
}

func getUserValidations(t *testing.T, db *sql.DB) {
	_, err := dal.GetUser(db, 0)
	AssertVError(t, err, verrors.InvalidRequestErrorCode, verrors.UserIDNegativeMessage)

	_, err = dal.GetUser(db, 999)
	AssertVError(t, err, verrors.NotFoundErrorCode, verrors.UserNotFoundMessage)
}

func updateEmailValidations(t *testing.T, db *sql.DB) {
	err := dal.UpdateUserEmail(nil, 0, "")
	AssertVError(t, err, verrors.UnexpectedErrorCode, verrors.DatabaseConnectionEmptyMessage)

	err = dal.UpdateUserEmail(db, 0, "")
	AssertVError(t, err, verrors.InvalidRequestErrorCode, verrors.UserIDNegativeMessage)

	err = dal.UpdateUserEmail(db, 1, "")
	AssertVError(t, err, verrors.InvalidRequestErrorCode, verrors.EmailEmptyMessage)

	err = dal.UpdateUserEmail(db, 1, "userwhisper.org")
	AssertVError(t, err, verrors.InvalidRequestErrorCode, "mail: missing '@' or angle-addr")

	err = dal.UpdateUserEmail(db, 1, "u@.org")
	AssertVError(t, err, verrors.InvalidRequestErrorCode, "mail: missing '@' or angle-addr")
}

func updateUserEmail(t *testing.T, db *sql.DB, user *models.User) *models.User {
	newMail := "user-modified@whisper.org"
	userID := user.ID
	err := dal.UpdateUserEmail(db, userID, newMail)
	AssertVErrorDoesNotExist(t, err)

	obtainedUser, err := dal.GetUser(db, userID)
	AssertVErrorDoesNotExist(t, err)

	Assert(t, obtainedUser.Email == newMail, "user mail mismatch.")
	return obtainedUser
}
