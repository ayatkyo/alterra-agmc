package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// 	Books
	e.GET("/books", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"message": "Get all books",
		})
	})

	e.GET("/books/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.JSON(http.StatusOK, map[string]any{
			"message": "Get book by id",
			"data":    id,
		})
	})

	e.POST("/books", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"message": "Create new book",
		})
	})

	e.PUT("/books/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.JSON(http.StatusOK, map[string]any{
			"message": "Update book by id",
			"data":    id,
		})
	})

	e.DELETE("/books/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.JSON(http.StatusOK, map[string]any{
			"message": "Delete book by id",
			"data":    id,
		})
	})

	return e
}
