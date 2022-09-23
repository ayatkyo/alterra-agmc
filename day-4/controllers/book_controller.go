package controllers

import (
	"fmt"
	"strconv"

	"github.com/ayatkyo/alterra-agcm/day-4/lib/database"
	"github.com/ayatkyo/alterra-agcm/day-4/models"
	"github.com/ayatkyo/alterra-agcm/day-4/utils"
	"github.com/labstack/echo/v4"
)

func BookGetAll(c echo.Context) error {
	return utils.ResponseSuccess(c, "Get all books", database.BookDB)
}

func BookGetByID(c echo.Context) error {
	// Parse id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.ResponseError(c, "Cannot parse book ID")
	}

	// Find book by id
	book, ok := utils.Find(database.BookDB, func(item models.Book) bool {
		return item.ID == uint(id)
	})
	if !ok {
		return utils.ResponseError(c, "Book not found")
	}

	return utils.ResponseSuccess(c, "Get book by ID", book)
}

func BookStore(c echo.Context) error {
	// Prepare
	book := &models.Book{}

	// Try bind data
	err := c.Bind(book)
	if err != nil {
		return utils.ResponseError(c, err.Error())
	}

	// validate
	if err = c.Validate(book); err != nil {
		return utils.ResponseError(c, err.Error())
	}

	// Create new book
	book.ID = uint(database.BookDBID) + 1
	database.BookDB = append(database.BookDB, *book)

	// Increment ID
	database.BookDBID = database.BookDBID + 1

	return utils.ResponseSuccess(c, "Book created", book)
}

func BookUpdate(c echo.Context) error {
	// Convert id param to integer
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.ResponseError(c, "Cannot parse book ID")
	}

	// Find book index
	idx, ok := utils.FindIndex(database.BookDB, func(item models.Book) bool {
		return item.ID == uint(id)
	})
	if !ok {
		return utils.ResponseError(c, "Book not found")
	}

	// validate
	req := models.Book{}
	err = c.Bind(&req)
	if err != nil {
		return utils.ResponseError(c, err.Error())
	}

	// validate
	if err = c.Validate(req); err != nil {
		return utils.ResponseError(c, err.Error())
	}

	// Update book
	book := &database.BookDB[idx]
	book.Title = req.Title
	book.Author = req.Author
	book.Year = req.Year

	return utils.ResponseSuccess(c, "Book updated", book)
}

func BookDestroy(c echo.Context) error {
	// Convert id param to integer
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.ResponseError(c, "Cannot parse book ID")
	}

	// Remove by id
	utils.Delete(&database.BookDB, func(item models.Book) bool {
		return item.ID == uint(id)
	})

	return utils.ResponseSuccess(c, fmt.Sprintf("Successfully delete book with ID %d", id), nil)
}
