package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"
	"time"

	"detailingpass/pkg/db"
	"detailingpass/web/templates/pages"

	"github.com/labstack/echo/v4"
)

var (
	bookingStatusOptions = []string{"pending", "confirmed", "declined", "cancelled"}
	bookingStatusSet     = map[string]bool{
		"pending":   true,
		"confirmed": true,
		"declined":  true,
		"cancelled": true,
	}
)

const adminBookingsPageSize int64 = 25

func (h *Handler) AdminBookings(c echo.Context) error {
	ctx := c.Request().Context()
	queries := db.New(h.db)

	page := parsePageParam(c.QueryParam("page"))
	offset := (int64(page) - 1) * adminBookingsPageSize

	rows, err := queries.ListBookings(ctx, db.ListBookingsParams{
		Limit:  adminBookingsPageSize,
		Offset: offset,
	})
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to load bookings")
	}

	total, _ := queries.CountBookings(ctx)
	pending, _ := queries.CountBookingsByStatus(ctx, sql.NullString{String: "pending", Valid: true})
	confirmed, _ := queries.CountBookingsByStatus(ctx, sql.NullString{String: "confirmed", Valid: true})
	declined, _ := queries.CountBookingsByStatus(ctx, sql.NullString{String: "declined", Valid: true})
	cancelled, _ := queries.CountBookingsByStatus(ctx, sql.NullString{String: "cancelled", Valid: true})

	items := make([]pages.AdminBookingItem, 0, len(rows))
	for _, row := range rows {
		items = append(items, buildAdminBookingItem(row))
	}

	hasNext := offset+int64(len(rows)) < total
	data := pages.AdminBookingsPageData{
		Stats: pages.AdminBookingStats{
			Total:     total,
			Pending:   pending,
			Confirmed: confirmed,
			Declined:  declined,
			Cancelled: cancelled,
		},
		Bookings:      items,
		StatusOptions: bookingStatusOptions,
		Pagination: pages.AdminPagination{
			Page:     page,
			PageSize: int(adminBookingsPageSize),
			Total:    total,
			HasPrev:  page > 1,
			HasNext:  hasNext,
			PrevPage: max(1, page-1),
			NextPage: page + 1,
		},
	}

	return pages.AdminBookings(data).Render(ctx, c.Response().Writer)
}

func (h *Handler) UpdateBookingStatus(c echo.Context) error {
	ctx := c.Request().Context()
	queries := db.New(h.db)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid booking ID")
	}

	status := normalizeBookingStatus(c.FormValue("status"))
	if !bookingStatusSet[status] {
		return c.String(http.StatusBadRequest, "Invalid booking status")
	}

	internalNotes := strings.TrimSpace(c.FormValue("internal_notes"))

	_, err = queries.UpdateBookingStatus(ctx, db.UpdateBookingStatusParams{
		Status: sql.NullString{
			String: status,
			Valid:  true,
		},
		InternalNotes: sql.NullString{
			String: internalNotes,
			Valid:  internalNotes != "",
		},
		ID: id,
	})
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to update booking")
	}

	redirect := "/admin/bookings"
	if pageParam := strings.TrimSpace(c.FormValue("page")); pageParam != "" {
		redirect = redirect + "?page=" + pageParam
	}
	return c.Redirect(http.StatusSeeOther, redirect)
}

func parsePageParam(raw string) int {
	page, err := strconv.Atoi(strings.TrimSpace(raw))
	if err != nil || page < 1 {
		return 1
	}
	return page
}

func normalizeBookingStatus(value string) string {
	value = strings.ToLower(strings.TrimSpace(value))
	if value == "" {
		return "pending"
	}
	return value
}

func buildAdminBookingItem(row db.Booking) pages.AdminBookingItem {
	startLocal := row.RequestedStart.In(bookingLocation)
	endLocal := row.RequestedEnd.In(bookingLocation)
	slotLabel, slotWindow := resolveSlotDetails(row.RequestedStart, row.RequestedEnd)

	var submittedAt string
	if row.CreatedAt.Valid {
		submittedAt = row.CreatedAt.Time.In(bookingLocation).Format("Jan 2, 2006 3:04 PM")
	}

	return pages.AdminBookingItem{
		ID:            row.ID,
		CustomerName:  row.CustomerName,
		Email:         row.Email,
		Phone:         nullableString(row.Phone),
		Service:       nullableString(row.ServiceInterest),
		Vehicle:       nullableString(row.VehicleDetails),
		Notes:         nullableString(row.Notes),
		Status:        normalizeBookingStatus(row.Status.String),
		SlotLabel:     slotLabel,
		SlotWindow:    slotWindow,
		DateLabel:     startLocal.Format("Monday, Jan 2"),
		SubmittedAt:   submittedAt,
		InternalNotes: nullableString(row.InternalNotes),
		Source:        nullableString(row.Source),
		StartISO:      startLocal.Format(time.RFC3339),
		EndISO:        endLocal.Format(time.RFC3339),
	}
}

func nullableString(value sql.NullString) string {
	if value.Valid {
		return value.String
	}
	return ""
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
