package book

import (
	"github.com/ayatkyo/alterra-agcm/day-6/internal/factory"
	"github.com/ayatkyo/alterra-agcm/day-6/pkg/utils"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewSevice(f),
	}
}

func (h *handler) GetAll(c echo.Context) error {
	books, err := h.service.FindAll(c.Request().Context())
	if err != nil {
		return utils.ResponseError(c, err.Error())
	}

	return utils.ResponseSuccess(c, "Get all books", books)
}

func (h *handler) GetByID(c echo.Context) error {
	return utils.ResponseError(c, "TODO")
}

func (h *handler) Store(c echo.Context) error {
	return utils.ResponseError(c, "TODO")
}

func (h *handler) Update(c echo.Context) error {
	return utils.ResponseError(c, "TODO")
}

func (h *handler) Destroy(c echo.Context) error {
	return utils.ResponseError(c, "TODO")
}
