package controllers

import (
	"strconv"

	"github.com/ayatkyo/alterra-agcm/day-3/lib/database"
	"github.com/ayatkyo/alterra-agcm/day-3/middlewares"
	"github.com/ayatkyo/alterra-agcm/day-3/models"
	"github.com/ayatkyo/alterra-agcm/day-3/utils"
	"github.com/labstack/echo/v4"
)

func UserGetAll(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		return utils.ResponseError(c, err.Error())
	}

	// hide some information
	list := []map[string]any{}
	for _, user := range users.([]models.User) {
		list = append(list, user.AsResponse())
	}

	return utils.ResponseSuccess(c, "Get all users", list)
}

func UserGetByID(c echo.Context) error {
	// Parse id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.ResponseError(c, "Cannot parse user ID")
	}

	// Find user
	user, err := database.GetUserByID(id)
	if err != nil {
		return utils.ResponseError(c, err.Error())
	}

	// Parse user
	userRes, ok := user.(models.User)
	if !ok {
		return utils.ResponseError(c, "Cannot parse user")
	}

	return utils.ResponseSuccess(c, "Get user by ID", userRes.AsResponse())
}

func UserStore(c echo.Context) error {
	// Get user data
	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		return utils.ResponseError(c, err.Error())
	}

	// Create password hash
	passwordHash, err := utils.BcryptMake(user.Password)
	if err != nil {
		return utils.ResponseError(c, err.Error())
	}

	// Set password hash
	user.Password = passwordHash

	// Create user
	newUser, err := database.CreateUser(user)
	if err != nil {
		return utils.ResponseError(c, err.Error())
	}

	// Parse user
	userRes, ok := newUser.(models.User)
	if !ok {
		return utils.ResponseError(c, "Cannot parse user")
	}

	return utils.ResponseSuccess(c, "User created", userRes.AsResponse())
}

func UserUpdate(c echo.Context) error {
	// Parse id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.ResponseError(c, "Cannot parse user ID")
	}

	// limit to only himself
	authID := middlewares.ExtractTokenUserID(c)
	if authID != id {
		return utils.ResponseError(c, "Cannot edit other user data")
	}

	// Get user input
	userData := models.User{}
	err = c.Bind(&userData)
	if err != nil {
		return utils.ResponseError(c, err.Error())
	}

	// Find user
	found, err := database.GetUserByID(id)
	if err != nil {
		return utils.ResponseError(c, err.Error())
	}

	// parse user
	user, ok := found.(models.User)
	if !ok {
		return utils.ResponseError(c, "Cannot parse user")
	}

	// Update data
	user.Fullname = userData.Fullname
	user.Email = userData.Email
	user.Username = userData.Username

	// Check password change
	if userData.Password != "" {
		// Create password hash
		passwordHash, err := utils.BcryptMake(userData.Password)
		if err != nil {
			return utils.ResponseError(c, err.Error())
		}

		// Set password hash
		user.Password = passwordHash
	}

	// Update in db
	_, err = database.UpdateUser(user)
	if err != nil {
		return utils.ResponseError(c, err.Error())
	}

	return utils.ResponseSuccess(c, "Your data has been updated", user.AsResponse())
}

func UserDestroy(c echo.Context) error {
	// Parse id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.ResponseError(c, "Cannot parse user ID")
	}

	// limit to only himself
	authID := middlewares.ExtractTokenUserID(c)
	if authID != id {
		return utils.ResponseError(c, "Cannot delete other user data")
	}

	// Delete user in db
	err = database.DeleteUser(id)
	if err != nil {
		return utils.ResponseError(c, err.Error())
	}

	return utils.ResponseSuccess(c, "Your account has been deleted", nil)
}
