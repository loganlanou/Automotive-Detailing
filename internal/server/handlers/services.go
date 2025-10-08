package handlers

import (
	"detailingpass/web/templates/pages"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Services(c echo.Context) error {
	// TODO: Fetch packages from database
	return pages.Services().Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) ServiceDetail(c echo.Context) error {
	slug := c.Param("slug")
	// TODO: Fetch package by slug from database
	return pages.ServiceDetail(slug).Render(c.Request().Context(), c.Response().Writer)
}
