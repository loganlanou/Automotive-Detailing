package handlers

import (
	"detailingpass/web/templates/pages"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Work(c echo.Context) error {
	// TODO: Parse filter params (make, model, year, package, price_min, price_max)
	// TODO: Fetch filtered jobs from database with pagination
	return pages.Work().Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) WorkDetail(c echo.Context) error {
	slug := c.Param("slug")
	// TODO: Fetch vehicle + job + media by slug
	return pages.WorkDetail(slug).Render(c.Request().Context(), c.Response().Writer)
}
