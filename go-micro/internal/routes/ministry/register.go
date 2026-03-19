package ministry

import (
	"github.com/labstack/echo/v5"
)

func RegisterRoutes(g *echo.Group, h *Handler) {
	g.GET("", h.GetMinistries)
	g.POST("/join", h.JoinMinistry)
}
