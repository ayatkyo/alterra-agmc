package routes

import (
	"net/http"
	"strconv"

	"github.com/ayatkyo/alterra-agcm/day-2/constant"
	"github.com/ayatkyo/alterra-agcm/day-2/models"
	"github.com/ayatkyo/alterra-agcm/day-2/utils"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// 	Books
	e.GET("/books", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"message": "Get all books",
			"data":    constant.StaticBookDB,
		})
	})

	e.GET("/books/:id", func(c echo.Context) error {
		//	Convert id param to integer
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(400, map[string]any{
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
