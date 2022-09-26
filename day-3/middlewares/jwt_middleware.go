package middlewares

import (
	"errors"
	"net/http"
	"time"

	"github.com/ayatkyo/alterra-agmc/day-3/constants"
	"github.com/ayatkyo/alterra-agmc/day-3/utils"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CreateToken(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userID"] = userID

	//	1 hour token
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constants.JWT_SECRET))
}

func ExtractTokenUserID(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if !user.Valid {
		return 0
	}

	claims := user.Claims.(jwt.MapClaims)
	return int(claims["userID"].(float64))
}

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
