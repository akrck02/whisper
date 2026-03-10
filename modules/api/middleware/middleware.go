// Package middleware provides functions to modify the context of a htttp request or response
package middleware

import (
	apimodels "github.com/akrck02/whisper/modules/api/models"
	verrors "github.com/akrck02/whisper/sdk/errors"
)

type Middleware func(context *apimodels.APIContext) *verrors.APIError
