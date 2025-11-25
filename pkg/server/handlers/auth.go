package handlers

import (
	"detailingpass/web/templates/pages"

	"github.com/labstack/echo/v4"
)

func (h *Handler) SignIn(c echo.Context) error {
	redirectUrl := c.QueryParam("redirect_url")
	if redirectUrl == "" {
		redirectUrl = "/"
	}
	return pages.SignIn(redirectUrl).Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) SignUp(c echo.Context) error {
	redirectUrl := c.QueryParam("redirect_url")
	if redirectUrl == "" {
		redirectUrl = "/"
	}
	return pages.SignUp(redirectUrl).Render(c.Request().Context(), c.Response().Writer)
}
