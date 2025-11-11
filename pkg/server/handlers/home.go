package handlers

import (
	"detailingpass/web/templates/pages"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Home(c echo.Context) error {
	// TODO: Fetch featured jobs from database
	return pages.Home().Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) About(c echo.Context) error {
	return pages.About().Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) FAQ(c echo.Context) error {
	return pages.FAQ().Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) Privacy(c echo.Context) error {
	return pages.Privacy().Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) Terms(c echo.Context) error {
	return pages.Terms().Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) StyleGuide(c echo.Context) error {
	return pages.StyleGuide().Render(c.Request().Context(), c.Response().Writer)
}
