package handlers

import (
	"detailingpass/web/templates/pages"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Reviews(c echo.Context) error {
	// TODO: Fetch reviews from database
	return pages.Reviews().Render(c.Request().Context(), c.Response().Writer)
}
