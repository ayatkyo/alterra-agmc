package user

import (
	"strconv"

	"github.com/ayatkyo/alterra-agcm/day-7/internal/dto"
	"github.com/ayatkyo/alterra-agcm/day-7/internal/factory"
	"github.com/ayatkyo/alterra-agcm/day-7/pkg/utils"
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

func (h *handler) GetAll(c echo.Context) error {
	users, err := h.service.FindAll(c.Request().Context())
	if err != nil {
		return utils.ResponseError(c, err.Error())
	}

	list := []map[string]any{}
	for _, user := range users {
		list = append(list, user.AsResponse())
	}

	return utils.ResponseSuccess(c, "Get all users", list)
}

func (h *handler) GetByID(c echo.Context) error {
	// Parse id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.ResponseError(c, "Cannot parse user ID")
	}

	// Find user
	user, err := h.service.FindByID(c.Request().Context(), uint(id))
	if err != nil {
		return utils.ResponseError(c, err.Error())
	}

	return utils.ResponseSuccess(c, "Get user by ID", user.AsResponse())
}

func (h *handler) Update(c echo.Context) error {
	// Parse id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.ResponseError(c, "Cannot parse user ID")
	}

	// limit to only himself
	authID := utils.JWTExtractUserID(c)
	if authID != id {
		return utils.ResponseError(c, "Cannot edit other user data")
	}

	payload := new(dto.UserUpdateRequest)
	if err := c.Bind(payload); err != nil {
		return utils.ResponseError(c, err.Error())
	}
	if err := c.Validate(payload); err != nil {
		return utils.ResponseError(c, err.Error())
	}

	user, err := h.service.Update(c.Request().Context(), uint(id), payload)
	if err != nil {
		return utils.ResponseError(c, err.Error())
	}

	return utils.ResponseSuccess(c, "Your data has been updated", user.AsResponse())
}

func (h *handler) Destroy(c echo.Context) error {
	// Parse id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.ResponseError(c, "Cannot parse user ID")
	}

	// limit to only himself
	authID := utils.JWTExtractUserID(c)
	if authID != id {
		return utils.ResponseError(c, "Cannot edit other user data")
	}

	err = h.service.Destroy(c.Request().Context(), uint(id))
	if err != nil {
		return utils.ResponseError(c, err.Error())
	}

	return utils.ResponseSuccess(c, "Your account has been deleted", nil)
}
