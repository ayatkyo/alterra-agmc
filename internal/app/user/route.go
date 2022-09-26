package user

import (
	"github.com/ayatkyo/alterra-agcm/day-10/internal/middlewares"
	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.Use(middlewares.JWTMiddleware)

	g.GET("", h.GetAll)
	g.GET("/:id", h.GetByID)
	g.PUT("/:id", h.Update)
	g.DELETE("/:id", h.Destroy)
}
