package sermon

import (
	"church-backend/internal/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) ListSermons(c *echo.Context) error {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	results, err := h.service.GetRecent(c.Request().Context(), limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to load sermons")
	}
	return c.JSON(http.StatusOK, results)
}

func (h *Handler) UploadSermon(c *echo.Context) error {
	req := new(models.Sermon)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid sermon data")
	}

	if err := h.service.repo.CreateSermon(c.Request().Context(), req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "could not save sermon")
	}
	return c.JSON(http.StatusCreated, req)
}

func RegisterRoutes(e *echo.Echo, h *Handler, adminMiddleware echo.MiddlewareFunc) {
	g := e.Group("/sermons")

	g.GET("", h.ListSermons)

	// only Admins can upload
	g.POST("", h.UploadSermon, adminMiddleware)
}
