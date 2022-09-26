package controllers

import (
	"github.com/ayatkyo/alterra-agmc/day-4/config"
	"github.com/ayatkyo/alterra-agmc/day-4/middlewares"
	"github.com/ayatkyo/alterra-agmc/day-4/models"
	"github.com/ayatkyo/alterra-agmc/day-4/utils"
	"github.com/labstack/echo/v4"
)

func AuthLogin(c echo.Context) error {
	req := models.UserLogin{}
	err := c.Bind(&req)
	if err != nil {
		return utils.ResponseError(c, err.Error())
	}

	if err = c.Validate(req); err != nil {
		return utils.ResponseError(c, err.Error())
	}

	// Find user
	var user models.User
	err = config.DB.Where("username = @username OR email = @username", map[string]any{"username": req.Username}).First(&user).Error
	if err != nil {
		return utils.ResponseError(c, err.Error())
	}

	// compare password
	if !utils.BcryptCheck(req.Password, user.Password) {
		return utils.ResponseError(c, "Username or password is incorrect")
	}

	// create token
	token, err := middlewares.CreateToken(int(user.ID))
	if err != nil {
		return utils.ResponseError(c, "Cannot create token")
	}

	return utils.ResponseSuccess(c, "Login success", token)
}
