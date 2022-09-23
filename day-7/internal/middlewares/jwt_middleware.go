package middlewares

import (
	"errors"
	"net/http"

	"github.com/ayatkyo/alterra-agcm/day-7/pkg/constants"
	"github.com/ayatkyo/alterra-agcm/day-7/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var JWTMiddleware echo.MiddlewareFunc = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(constants.JWT_SECRET),
	ErrorHandlerWithContext: func(err error, c echo.Context) error {
		var message string
		var code int

		if errors.Is(err, middleware.ErrJWTMissing) {
			message = "Missing or malformed JWT"
			code = http.StatusBadRequest
		} else if errors.Is(err, middleware.ErrJWTInvalid) {
			message = "Invalid or expired jwt"
			code = http.StatusUnauthorized
		} else {
			message = err.Error()
			code = http.StatusBadRequest
		}

		return utils.ResponseErrorWithCode(c, message, code)
	},
})
