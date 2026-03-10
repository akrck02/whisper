package middleware

import (
	"github.com/akrck02/whisper/database"
	apimodels "github.com/akrck02/whisper/modules/api/models"
	verrors "github.com/akrck02/whisper/sdk/errors"
)

func Database(context *apimodels.APIContext) *verrors.APIError {
	if !context.Trazability.Endpoint.Database || nil != context.Database {
		return nil
	}

	db, err := database.GetConnection()
	if nil != err {
		return verrors.NewAPIError(verrors.DatabaseError(verrors.CannotConnectToDatabaseMessage))
	}

	context.Database = db
	if nil == context.Database {
		return verrors.NewAPIError(verrors.DatabaseError(verrors.CannotConnectToDatabaseMessage))
	}

	return nil
}
