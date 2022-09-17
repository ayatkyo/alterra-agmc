package routes

import (
	"github.com/ayatkyo/alterra-agcm/day-2/controllers"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// Books
	e.GET("/books", controllers.BookGetAll)
	e.GET("/books/:id", controllers.BookGetByID)
	e.POST("/books", controllers.BookStore)
	e.PUT("/books/:id", controllers.BookUpdate)
	e.DELETE("/books/:id", controllers.BookDestroy)

	// User
	e.GET("/users", controllers.UserGetAll)
	e.GET("/users/:id", controllers.UserGetByID)
	e.POST("/users", controllers.UserStore)
	e.PUT("/users/:id", controllers.UserUpdate)
	e.DELETE("/users/:id", controllers.UserDestroy)

	return e
}
