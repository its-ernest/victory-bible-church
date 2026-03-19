package ministry

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

// GetMinistries returns the list for the dropdown/grid
func (h *Handler) GetMinistries(c *echo.Context) error {
	list, err := h.service.ListMinistries(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch ministries")
	}
	return c.JSON(http.StatusOK, list)
}

// JoinMinistry handles the POST request
func (h *Handler) JoinMinistry(c *echo.Context) error {
	// extract member phone from jwt
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*jwt.RegisteredClaims)
	phone := claims.Subject

	var req struct {
		MinistryID string `json:"ministry_id"`
	}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid ministry ID")
	}

	err := h.service.Join(c.Request().Context(), phone, req.MinistryID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "could not join ministry")
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "joined successfully"})
}
