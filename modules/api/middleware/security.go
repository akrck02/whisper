package middleware

import (
	"github.com/akrck02/whisper/database"
	apimodels "github.com/akrck02/whisper/modules/api/models"
	verrors "github.com/akrck02/whisper/sdk/errors"
)

const AuthorizationHeader = "Authorization"

func Security(context *apimodels.APIContext) *verrors.APIError {

	//TODO temporary
	db, err := database.GetConnection()
	context.Database = db
	return nil
	// Check if endpoint is registered and secured
	if !context.Trazability.Endpoint.Secured {
		return nil
	}

	// Check if token is empty
	if context.Request.Authorization == "" {
		return verrors.NewAPIError(verrors.Unauthorized(verrors.TokenEmptyMessage))
	}

	db, err = database.GetConnection()
	if nil != err {
		return verrors.NewAPIError(verrors.Unauthorized(verrors.TokenInvalidMessage))
	}
	context.Database = db

	// Check if token is valid
	//	token, _ := strings.CutPrefix(context.Request.Authorization, "Bearer ")
	//	_, verr := dal.LoginWithAuth(context.Database, context.Configuration.JWTSecret, token)
	//	if nil != verr {
	//		logger.Error(token)
	//		logger.Error(verr)
	//		return verrors.NewAPIError(verrors.Unauthorized(verrors.TokenInvalidMessage))
	//	}

	return nil
}
