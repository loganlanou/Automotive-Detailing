package handlers

import (
	"detailingpass/web/templates/pages"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Blog(c echo.Context) error {
	// TODO: Fetch published posts with pagination
	return pages.Blog().Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) BlogPost(c echo.Context) error {
	slug := c.Param("slug")
	// TODO: Fetch post by slug
	return pages.BlogPost(slug).Render(c.Request().Context(), c.Response().Writer)
}
