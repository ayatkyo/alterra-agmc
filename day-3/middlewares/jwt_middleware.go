package middlewares

import (
	"time"

	"github.com/ayatkyo/alterra-agcm/day-3/constants"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
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
	return claims["userID"].(int)
}
