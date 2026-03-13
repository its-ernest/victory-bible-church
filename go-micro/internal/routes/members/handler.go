package members

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

// GetProfile handles GET /members/profile
func (h *Handler) GetProfile(c *echo.Context) error {
	// extract phone from jwt(Subject claim)
	val := c.Get("user")
    if val == nil {
        return echo.NewHTTPError(http.StatusUnauthorized, "Token not found in context")
    }
    token, ok := val.(*jwt.Token)
    if !ok {
        return echo.NewHTTPError(http.StatusInternalServerError, "Invalid token type in context")
    }
	claims := token.Claims.(*jwt.RegisteredClaims)
	phone := claims.Subject

	member, err := h.service.GetProfile(c.Request().Context(), phone)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Error: failed to get user: "+ err.Error())
	}
	return c.JSON(http.StatusOK, member)
}

// UpdateRequest defines the allowed payload for updating a profile
type UpdateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// UpdateProfile handles PUT /members/profile
func (h *Handler) UpdateProfile(c *echo.Context) error {
	// extract identity from jwt
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*jwt.RegisteredClaims)
	phone := claims.Subject

	// json body
	req := new(UpdateRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request format")
	}

	err := h.service.UpdateMember(
		c.Request().Context(),
		phone,
		req.FirstName,
		req.LastName,
		req.Email,
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "profile updated successfully",
	})
}

// RegisterRoutes ties the handlers to the Echo group
func RegisterRoutes(g *echo.Group, h *Handler) {
	g.GET("/profile", h.GetProfile)
	g.PUT("/profile", h.UpdateProfile)
}