package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ResponseError(c echo.Context, message string) error {
	return c.JSON(http.StatusBadRequest, map[string]any{
		"success": false,
		"message": message,
	})
}

func ResponseSuccess(c echo.Context, message string, data any) error {
	res := map[string]any{
		"success": true,
		"message": message,
	}

	if data != nil {
		res["data"] = data
	}

	return c.JSON(http.StatusOK, res)
}
