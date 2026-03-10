package api

import (
	"github.com/akrck02/whisper/modules/api/controllers"
	apimodels "github.com/akrck02/whisper/modules/api/models"
)

var EndpointRegistry = []apimodels.Endpoint{
	controllers.UserCreateEndpoint,
	controllers.UserGetEndpoint,
	controllers.UserGetByEmailEndpoint,
}
