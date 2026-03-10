// Package userservice provides functions to handle the api http request related to a user
package controllers

import (
	"io"
	"net/http"

	"github.com/akrck02/whisper/database/dal"
	apimodels "github.com/akrck02/whisper/modules/api/models"
	services "github.com/akrck02/whisper/modules/api/services"
	verrors "github.com/akrck02/whisper/sdk/errors"
	inout "github.com/akrck02/whisper/sdk/io"
	"github.com/akrck02/whisper/sdk/models"
)

// USER REGISTER
var UserCreateEndpoint = apimodels.Endpoint{
	Path:     "users",
	Method:   apimodels.PostMethod,
	Listener: register,
	Secured:  false,
	Database: true,
}

func register(context *apimodels.APIContext) (*apimodels.Response, *verrors.APIError) {
	body := context.Request.Body.(io.ReadCloser)
	user := models.User{}
	err := inout.ParseJSON(&body, &user)
	if nil != err {
		return nil, verrors.NewAPIError(verrors.New(verrors.InvalidRequestErrorCode, "Request body has invalid format."))
	}

	context.Request.Body = user

	usr := context.Request.Body.(models.User)
	userUuid, rerr := dal.CreateUser(context.Database, &usr)

	if nil != rerr {
		return nil, verrors.NewAPIError(rerr)
	}

	return &apimodels.Response{
		Code:     http.StatusOK,
		Response: userUuid,
	}, nil
}

var UserGetEndpoint = apimodels.Endpoint{
	Path:     "users/{uuid}",
	Method:   apimodels.GetMethod,
	Listener: get,
	Secured:  true,
}

func get(context *apimodels.APIContext) (*apimodels.Response, *verrors.APIError) {
	uuid := context.Request.Params["uuid"]
	if uuid == "" {
		return nil, verrors.NewAPIError(verrors.New(verrors.InvalidRequestErrorCode, "user uuid cannot be empty"))
	}

	user, getUserErr := services.GetUser(context.Database, uuid)
	if nil != getUserErr {
		return nil, verrors.NewAPIError(getUserErr)
	}

	return &apimodels.Response{
		Code:     http.StatusOK,
		Response: user,
	}, nil
}

var UserGetByEmailEndpoint = apimodels.Endpoint{
	Path:     "users/email",
	Method:   apimodels.GetMethod,
	Listener: getByEmail,
	Secured:  true,
}

func getByEmail(context *apimodels.APIContext) (*apimodels.Response, *verrors.APIError) {
	email := context.Request.Params["email"]
	user, err := services.GetUserByEmail(context.Database, email)
	if nil != err {
		return nil, verrors.NewAPIError(err)
	}

	return &apimodels.Response{
		Code:     http.StatusOK,
		Response: user,
	}, nil
}
