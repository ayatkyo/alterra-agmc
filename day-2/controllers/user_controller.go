package controllers

import (
	"net/http"
	"strconv"

	"github.com/ayatkyo/alterra-agcm/day-2/lib/database"
	"github.com/ayatkyo/alterra-agcm/day-2/models"
	"github.com/ayatkyo/alterra-agcm/day-2/utils"
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
	//	Parse id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "Cannot parse id",
			"success": false,
		})
	}

	//	Find user
	user, err := database.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
			"success": false,
		})
	}

	userRes, ok := user.(models.User)
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "Cannot parse user",
			"success": false,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "Get user by id",
		"data": map[string]any{
			"id":         userRes.ID,
			"fullname":   userRes.Fullname,
			"username":   userRes.Username,
			"email":      userRes.Email,
			"created_at": userRes.CreatedAt,
			"updated_at": userRes.UpdatedAt,
		},
	})
}

func UserStore(c echo.Context) error {
	//	Get user data
	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//	Create password hash
	passwordHash, err := utils.BcryptMake(user.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//	Set password hash
	user.Password = passwordHash

	//	Create user
	newUser, err := database.CreateUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userRes, ok := newUser.(models.User)
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "Cannot parse user",
			"success": false,
		})
	}

	//	Return result
	return c.JSON(http.StatusOK, map[string]any{
		"message": "User created",
		"data": map[string]any{
			"id":         userRes.ID,
			"fullname":   userRes.Fullname,
			"username":   userRes.Username,
			"email":      userRes.Email,
			"created_at": userRes.CreatedAt,
			"updated_at": userRes.UpdatedAt,
		},
	})
}

func UserUpdate(c echo.Context) error {
	return c.String(200, "Update user by id")
}

func UserDestroy(c echo.Context) error {
	return c.String(200, "Delete user by id")
}
