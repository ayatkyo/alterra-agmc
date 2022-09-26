package book

import (
	"fmt"
	"strconv"

	"github.com/ayatkyo/alterra-agcm/day-10/internal/factory"
	"github.com/ayatkyo/alterra-agcm/day-10/internal/models"
	"github.com/ayatkyo/alterra-agcm/day-10/pkg/utils"
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
	// Parse id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.ResponseError(c, "cnnot parse book ID")
	}

	book, err := h.service.FindByID(c.Request().Context(), uint(id))
	if err != nil {
		return utils.ResponseError(c, err.Error())
	}

	return utils.ResponseSuccess(c, "Get book by ID", book)
}

func (h *handler) Store(c echo.Context) error {
	payload := new(models.Book)
	if err := c.Bind(payload); err != nil {
		return utils.ResponseError(c, err.Error())
	}
	if err := c.Validate(payload); err != nil {
		return utils.ResponseError(c, err.Error())
	}

	book, err := h.service.Store(c.Request().Context(), payload)
	if err != nil {
		return utils.ResponseError(c, err.Error())
	}

	return utils.ResponseSuccess(c, "Book created", book)
}

func (h *handler) Update(c echo.Context) error {
	// Parse id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.ResponseError(c, "cnnot parse book ID")
	}

	payload := new(models.Book)
	if err := c.Bind(payload); err != nil {
		return utils.ResponseError(c, err.Error())
	}
	if err := c.Validate(payload); err != nil {
		return utils.ResponseError(c, err.Error())
	}

	book, err := h.service.Update(c.Request().Context(), uint(id), payload)
	if err != nil {
		return utils.ResponseError(c, err.Error())
	}

	return utils.ResponseSuccess(c, "Book updated", book)
}

func (h *handler) Destroy(c echo.Context) error {
	// Parse id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.ResponseError(c, "cnnot parse book ID")
	}

	err = h.service.Destroy(c.Request().Context(), uint(id))
	if err != nil {
		return utils.ResponseError(c, err.Error())
	}

	return utils.ResponseSuccess(c, fmt.Sprintf("Successfully delete book with ID %d", id), nil)
}
