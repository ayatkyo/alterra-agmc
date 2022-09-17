package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserGetAll(c echo.Context) error {
	return c.String(http.StatusOK, "Get all users")
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
