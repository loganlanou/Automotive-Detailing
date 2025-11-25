package handlers

import (
	"fmt"
	"time"

	"detailingpass/web/templates/pages"

	"github.com/labstack/echo/v4"
)

func (h *Handler) BookingPage(c echo.Context) error {
	var slots []pages.BookingSlot
	for _, slot := range bookingSlotDefinitions {
		slots = append(slots, pages.BookingSlot{
			ID:          slot.ID,
			Label:       slot.Label,
			Description: slot.Description,
			Duration:    formatSlotDuration(slot.Duration),
		})
	}

	data := pages.BookingPageData{
		Slots: slots,
	}

	return pages.Booking(data).Render(c.Request().Context(), c.Response().Writer)
}

func formatSlotDuration(d time.Duration) string {
	hours := d.Hours()
	if hours == float64(int(hours)) {
		return fmt.Sprintf("%.0f hrs", hours)
	}
	return fmt.Sprintf("%.1f hrs", hours)
}
