package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ayatkyo/alterra-agmc/day-2/constant"
	"github.com/ayatkyo/alterra-agmc/day-2/models"
	"github.com/ayatkyo/alterra-agmc/day-2/utils"
	"github.com/labstack/echo/v4"
)

func BookGetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"message": "Get all books",
		"data":    constant.StaticBookDB,
	})
}

func BookGetByID(c echo.Context) error {
	//	Convert id param to integer
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "Cannot convert id to int",
			"success": false,
		})
	}

	//	Find book by id
	book, ok := utils.Find(constant.StaticBookDB, func(item models.Book) bool {
		return item.ID == uint(id)
	})

	//	Book not found
	if !ok {
		return c.JSON(400, map[string]any{
			"message": "Book not found",
			"success": false,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"success": true,
		"message": "Get book by id",
		"data":    book,
	})
}

func BookStore(c echo.Context) error {
	//	Prepare
	book := &models.Book{
		ID: uint(constant.StaticBookDBID) + 1,
	}

	//	Try bind data
	err := c.Bind(book)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"success": false,
			"message": err.Error(),
		})
	}

	//	Create new book
	constant.StaticBookDB = append(constant.StaticBookDB, *book)

	//	Increment ID
	constant.StaticBookDBID = constant.StaticBookDBID + 1

	return c.JSON(http.StatusOK, map[string]any{
		"message": "Create new book",
		"data":    book,
	})
}

func BookUpdate(c echo.Context) error {
	//	Convert id param to integer
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "Cannot convert id to int",
			"success": false,
		})
	}

	//	Find book
	var idx int
	ok := false
	for i, item := range constant.StaticBookDB {
		idx = i
		if item.ID == uint(id) {
			ok = true
			break
		}
	}
	if !ok {
		return c.JSON(400, map[string]any{
			"message": "Book not found",
			"success": false,
		})
	}

	//	Update book
	book := &constant.StaticBookDB[idx]
	c.Bind(book)

	return c.JSON(http.StatusOK, map[string]any{
		"message": "Update book by id",
		"data":    book,
	})
}

func BookDestroy(c echo.Context) error {
	//	Convert id param to integer
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "Cannot convert id to int",
			"success": false,
		})
	}

	//	Remove by id
	utils.Delete(&constant.StaticBookDB, func(item models.Book) bool {
		return item.ID == uint(id)
	})

	return c.JSON(http.StatusOK, map[string]any{
		"success": true,
		"message": fmt.Sprintf("Successfully delete book with id %d", id),
	})
}
