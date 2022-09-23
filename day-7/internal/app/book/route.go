package book

import (
	"github.com/ayatkyo/alterra-agcm/day-7/internal/middlewares"
	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.GetAll)
	g.GET("/:id", h.GetByID)
	g.POST("", h.Store, middlewares.JWTMiddleware)
	g.PUT("/:id", h.Update, middlewares.JWTMiddleware)
	g.DELETE("/:id", h.Destroy, middlewares.JWTMiddleware)
}
