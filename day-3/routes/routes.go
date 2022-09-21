package routes

import (
	"github.com/ayatkyo/alterra-agcm/day-3/controllers"
	"github.com/ayatkyo/alterra-agcm/day-3/middlewares"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}

	return nil
}

func New() *echo.Echo {
	e := echo.New()

	// Validator
	e.Validator = &CustomValidator{validator: validator.New()}

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
