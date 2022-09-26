package auth

import (
	"github.com/ayatkyo/alterra-agcm/day-10/internal/dto"
	"github.com/ayatkyo/alterra-agcm/day-10/internal/factory"
	"github.com/ayatkyo/alterra-agcm/day-10/pkg/utils"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

func (h *handler) Login(c echo.Context) error {
	payload := new(dto.AuthLoginRequest)

	if err := c.Bind(payload); err != nil {
		return utils.ResponseError(c, err.Error())
	}
	if err := c.Validate(payload); err != nil {
		return utils.ResponseError(c, err.Error())
	}

	res, err := h.service.Login(c.Request().Context(), payload)
	if err != nil {
		return utils.ResponseError(c, err.Error())
	}

	return utils.ResponseSuccess(c, "Login success", res)
}

func (h *handler) Signup(c echo.Context) error {
	payload := new(dto.AuthSignupRequest)

	if err := c.Bind(payload); err != nil {
		return utils.ResponseError(c, err.Error())
	}
	if err := c.Validate(payload); err != nil {
		return utils.ResponseError(c, err.Error())
	}

	res, err := h.service.Signup(c.Request().Context(), payload)
	if err != nil {
		return utils.ResponseError(c, err.Error())
	}

	return utils.ResponseSuccess(c, "Signup success", res)
}
