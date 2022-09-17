package controllers

import (
	"net/http"

	"github.com/ayatkyo/alterra-agcm/day-2/lib/database"
	"github.com/labstack/echo/v4"
)

func UserGetAll(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "Get all users",
		"data":    users,
	})
}

func UserGetByID(c echo.Context) error {
	return c.String(200, "Get user by id")
}

func UserStore(c echo.Context) error {
	return c.String(200, "Create new user")
}

func UserUpdate(c echo.Context) error {
	return c.String(200, "Update user by id")
}

func UserDestroy(c echo.Context) error {
	return c.String(200, "Delete user by id")
}
