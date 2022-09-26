package http

import (
	"github.com/ayatkyo/alterra-agcm/day-10/internal/app/auth"
	"github.com/ayatkyo/alterra-agcm/day-10/internal/app/book"
	"github.com/ayatkyo/alterra-agcm/day-10/internal/app/user"
	"github.com/ayatkyo/alterra-agcm/day-10/internal/factory"
	"github.com/ayatkyo/alterra-agcm/day-10/pkg/utils"
	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	e.Validator = utils.EchoCustomValidator

	e.GET("/", func(c echo.Context) error {
		return utils.ResponseSuccess(c, "Alterra Day 10 (CI/CD) - Hidayatullah", nil)
	})

	e.GET("/status", func(c echo.Context) error {
		return utils.ResponseSuccess(c, "Alterra Day 10 (CI/CD) - API Status", "OK")
	})

	api := e.Group("/api")
	auth.NewHandler(f).Route(api.Group("/auth"))
	user.NewHandler(f).Route(api.Group("/users"))
	book.NewHandler(f).Route(api.Group("/books"))
}
