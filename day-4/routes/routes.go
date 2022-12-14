package routes

import (
	"github.com/ayatkyo/alterra-agmc/day-4/controllers"
	"github.com/ayatkyo/alterra-agmc/day-4/middlewares"
	"github.com/ayatkyo/alterra-agmc/day-4/utils"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// Validator
	e.Validator = utils.EchoCustomValidator

	// Auth
	e.POST("/login", controllers.AuthLogin)

	// Books
	e.GET("/books", controllers.BookGetAll)
	e.GET("/books/:id", controllers.BookGetByID)
	e.POST("/books", controllers.BookStore, middlewares.JWTMiddleware)
	e.PUT("/books/:id", controllers.BookUpdate, middlewares.JWTMiddleware)
	e.DELETE("/books/:id", controllers.BookDestroy, middlewares.JWTMiddleware)

	// User
	e.GET("/users", controllers.UserGetAll, middlewares.JWTMiddleware)
	e.GET("/users/:id", controllers.UserGetByID, middlewares.JWTMiddleware)
	e.POST("/users", controllers.UserStore)
	e.PUT("/users/:id", controllers.UserUpdate, middlewares.JWTMiddleware)
	e.DELETE("/users/:id", controllers.UserDestroy, middlewares.JWTMiddleware)

	return e
}
