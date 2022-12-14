package http

import (
	"github.com/ayatkyo/alterra-agmc/day-6/internal/app/auth"
	"github.com/ayatkyo/alterra-agmc/day-6/internal/app/book"
	"github.com/ayatkyo/alterra-agmc/day-6/internal/app/user"
	"github.com/ayatkyo/alterra-agmc/day-6/internal/factory"
	"github.com/ayatkyo/alterra-agmc/day-6/pkg/utils"
	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	e.Validator = utils.EchoCustomValidator

	e.GET("/status", func(c echo.Context) error {
		return utils.ResponseSuccess(c, "Alterra Day 6 API Status", "OK")
	})

	api := e.Group("/api")
	auth.NewHandler(f).Route(api.Group("/auth"))
	user.NewHandler(f).Route(api.Group("/users"))
	book.NewHandler(f).Route(api.Group("/books"))
}
