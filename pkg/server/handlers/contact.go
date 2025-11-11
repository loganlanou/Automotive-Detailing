package handlers

import (
	"detailingpass/web/templates/pages"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Contact(c echo.Context) error {
	return pages.Contact().Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) ContactSubmit(c echo.Context) error {
	// TODO: Parse form data
	// TODO: Validate (spam trap, required fields)
	// TODO: Send email via SMTP
	// TODO: Return success/error response

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Form submitted successfully",
	})
}
