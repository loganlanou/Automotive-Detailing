package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"detailingpass/pkg/auth"
	"detailingpass/pkg/db"

	"github.com/labstack/echo/v4"
)

const (
	defaultBookingTimezone = "America/New_York"
	defaultBookingHorizon  = 45
	maxBookingHorizon      = 90
	defaultBookingDays     = 30
)

var (
	bookingLocation        = loadBookingLocation()
	bookingSlotDefinitions = []slotDefinition{
		{
			ID:          "morning-detail",
			Label:       "Morning Detail",
			Description: "Kick off the day with a full refresh.",
			StartHour:   8,
			StartMinute: 0,
			Duration:    3 * time.Hour,
		},
		{
			ID:          "midday-refresh",
			Label:       "Midday Refresh",
			Description: "Great for exterior + interior combos.",
			StartHour:   12,
			StartMinute: 30,
			Duration:    3 * time.Hour,
		},
		{
			ID:          "late-day-polish",
			Label:       "Late Day Polish",
			Description: "Perfect for after-work drop-offs.",
			StartHour:   16,
			StartMinute: 0,
			Duration:    3 * time.Hour,
		},
	}
	slotLookup = buildSlotLookup(bookingSlotDefinitions)
	// All days are available for booking
	closedDays = map[time.Weekday]bool{}
)

type slotDefinition struct {
	ID          string
	Label       string
	Description string
	StartHour   int
	StartMinute int
	Duration    time.Duration
}

type availabilityResponse struct {
	GeneratedAt time.Time         `json:"generated_at"`
	Range       availabilityRange `json:"range"`
	Days        []availabilityDay `json:"days"`
	Slots       []slotMeta        `json:"slots"`
}

type availabilityRange struct {
	Start string `json:"start"`
	End   string `json:"end"`
	Days  int    `json:"days"`
}

type availabilityDay struct {
	Date            string             `json:"date"`
	Label           string             `json:"label"`
	IsToday         bool               `json:"is_today"`
	IsWeekend       bool               `json:"is_weekend"`
	IsClosed        bool               `json:"is_closed"`
	HasAvailability bool               `json:"has_availability"`
	Slots           []slotAvailability `json:"slots"`
}

type slotMeta struct {
	ID          string `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Duration    string `json:"duration"`
}

type slotAvailability struct {
	ID        string `json:"id"`
	Label     string `json:"label"`
	Window    string `json:"window"`
	StartISO  string `json:"start_iso"`
	EndISO    string `json:"end_iso"`
	Available bool   `json:"available"`
}

type bookingRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Vehicle string `json:"vehicle"`
	Service string `json:"service"`
	Notes   string `json:"notes"`
	Date    string `json:"date"`
	SlotID  string `json:"slot_id"`
}

type bookingResponse struct {
	Message string                 `json:"message"`
	Booking map[string]interface{} `json:"booking"`
}

func loadBookingLocation() *time.Location {
	tz := os.Getenv("BOOKING_TIMEZONE")
	if tz == "" {
		tz = defaultBookingTimezone
	}
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return time.Local
	}
	return loc
}

func buildSlotLookup(slots []slotDefinition) map[string]slotDefinition {
	out := make(map[string]slotDefinition, len(slots))
	for _, slot := range slots {
		out[slot.ID] = slot
	}
	return out
}

func (h *Handler) BookingAvailability(c echo.Context) error {
	ctx := c.Request().Context()
	queries := db.New(h.db)

	start := time.Now().In(bookingLocation).Truncate(24 * time.Hour)
	if monthParam := strings.TrimSpace(c.QueryParam("month")); monthParam != "" {
		if parsed, err := time.ParseInLocation("2006-01", monthParam, bookingLocation); err == nil {
			firstOfMonth := time.Date(parsed.Year(), parsed.Month(), 1, 0, 0, 0, 0, bookingLocation)
			if firstOfMonth.After(start) {
				start = firstOfMonth
			}
		}
	}

	if startParam := strings.TrimSpace(c.QueryParam("start")); startParam != "" {
		if parsed, err := time.ParseInLocation("2006-01-02", startParam, bookingLocation); err == nil {
			if parsed.After(start) {
				start = parsed
			} else if parsed.Equal(start) {
				start = parsed
			}
		}
	}

	daysRequested := defaultBookingDays
	if daysParam := strings.TrimSpace(c.QueryParam("days")); daysParam != "" {
		if parsed, err := strconv.Atoi(daysParam); err == nil && parsed > 0 {
			if parsed > maxBookingHorizon {
				parsed = maxBookingHorizon
			}
			daysRequested = parsed
		}
	}

	if daysRequested > defaultBookingHorizon {
		daysRequested = defaultBookingHorizon
	}

	endExclusive := start.AddDate(0, 0, daysRequested)
	blocked, err := queries.ListBlockedSlots(ctx, db.ListBlockedSlotsParams{
		RequestedStart:   start.UTC(),
		RequestedStart_2: endExclusive.UTC(),
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Unable to load availability",
		})
	}

	blockedMap := make(map[string]struct{}, len(blocked))
	for _, slot := range blocked {
		blockedMap[slotKey(slot.RequestedStart)] = struct{}{}
	}

	days := buildAvailabilityDays(start, endExclusive, blockedMap)
	resp := availabilityResponse{
		GeneratedAt: time.Now().In(bookingLocation),
		Range: availabilityRange{
			Start: start.Format("2006-01-02"),
			End:   endExclusive.AddDate(0, 0, -1).Format("2006-01-02"),
			Days:  daysRequested,
		},
		Days:  days,
		Slots: buildSlotMeta(),
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) CreateBookingRequest(c echo.Context) error {
	var req bookingRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))
	req.Phone = strings.TrimSpace(req.Phone)

	if req.Name == "" || req.Email == "" || req.Date == "" || req.SlotID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Name, email, date, and slot are required"})
	}

	slotDef, ok := slotLookup[req.SlotID]
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid slot selection"})
	}

	day, err := time.ParseInLocation("2006-01-02", req.Date, bookingLocation)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid date format"})
	}

	slotStartLocal := time.Date(day.Year(), day.Month(), day.Day(), slotDef.StartHour, slotDef.StartMinute, 0, 0, bookingLocation)
	if slotStartLocal.Before(time.Now().In(bookingLocation)) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Selected slot is no longer in the future"})
	}

	if closedDays[slotStartLocal.Weekday()] {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "We are closed on the selected date"})
	}

	if slotStartLocal.After(time.Now().In(bookingLocation).AddDate(0, 0, maxBookingHorizon)) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Selected slot is outside our booking window"})
	}

	slotStartUTC := slotStartLocal.UTC()
	slotEndUTC := slotStartUTC.Add(slotDef.Duration)

	ctx := c.Request().Context()
	queries := db.New(h.db)

	conflictCount, err := queries.CountBlockedSlotsAt(ctx, slotStartUTC)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to save booking"})
	}
	if conflictCount > 0 {
		return c.JSON(http.StatusConflict, map[string]string{"error": "That time has just been taken. Choose a different slot."})
	}

	// Get Clerk user ID from session if logged in
	clerkUserID := auth.GetUserID(ctx)

	booking, err := queries.CreateBooking(ctx, db.CreateBookingParams{
		CustomerName: req.Name,
		Email:        req.Email,
		Phone: sql.NullString{
			String: req.Phone,
			Valid:  req.Phone != "",
		},
		VehicleDetails: sql.NullString{
			String: strings.TrimSpace(req.Vehicle),
			Valid:  strings.TrimSpace(req.Vehicle) != "",
		},
		ServiceInterest: sql.NullString{
			String: strings.TrimSpace(req.Service),
			Valid:  strings.TrimSpace(req.Service) != "",
		},
		Notes: sql.NullString{
			String: strings.TrimSpace(req.Notes),
			Valid:  strings.TrimSpace(req.Notes) != "",
		},
		RequestedStart: slotStartUTC,
		RequestedEnd:   slotEndUTC,
		Status: sql.NullString{
			String: "pending",
			Valid:  true,
		},
		Source: sql.NullString{
			String: "web",
			Valid:  true,
		},
		ClerkUserID: sql.NullString{
			String: clerkUserID,
			Valid:  clerkUserID != "",
		},
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to save booking"})
	}

	resp := bookingResponse{
		Message: "Booking request received. We'll confirm shortly.",
		Booking: map[string]interface{}{
			"id":          booking.ID,
			"name":        booking.CustomerName,
			"email":       booking.Email,
			"status":      booking.Status.String,
			"slot_label":  slotDef.Label,
			"slot_window": fmt.Sprintf("%s – %s", slotStartLocal.Format("3:04 PM"), slotStartLocal.Add(slotDef.Duration).Format("3:04 PM")),
			"date":        slotStartLocal.Format("Monday, January 2"),
		},
	}

	return c.JSON(http.StatusCreated, resp)
}

func buildAvailabilityDays(start, endExclusive time.Time, blocked map[string]struct{}) []availabilityDay {
	var days []availabilityDay
	now := time.Now().In(bookingLocation)
	currentDay := start

	for currentDay.Before(endExclusive) {
		startOfDay := time.Date(currentDay.Year(), currentDay.Month(), currentDay.Day(), 0, 0, 0, 0, bookingLocation)
		isClosed := closedDays[startOfDay.Weekday()]
		slots := make([]slotAvailability, 0, len(bookingSlotDefinitions))
		hasAvailability := false

		if !isClosed {
			for _, slot := range bookingSlotDefinitions {
				slotStartLocal := time.Date(startOfDay.Year(), startOfDay.Month(), startOfDay.Day(), slot.StartHour, slot.StartMinute, 0, 0, bookingLocation)
				slotEndLocal := slotStartLocal.Add(slot.Duration)
				slotStartUTC := slotStartLocal.UTC()

				available := slotStartLocal.After(now)
				if _, exists := blocked[slotKey(slotStartUTC)]; exists {
					available = false
				}

				windowLabel := slotWindowLabel(slotStartLocal, slot.Duration)
				slots = append(slots, slotAvailability{
					ID:        slot.ID,
					Label:     slot.Label,
					Window:    windowLabel,
					StartISO:  slotStartLocal.Format(time.RFC3339),
					EndISO:    slotEndLocal.Format(time.RFC3339),
					Available: available,
				})

				if available {
					hasAvailability = true
				}
			}
		}

		days = append(days, availabilityDay{
			Date:            startOfDay.Format("2006-01-02"),
			Label:           startOfDay.Format("Mon, Jan 2"),
			IsToday:         startOfDay.Equal(now.Truncate(24 * time.Hour)),
			IsWeekend:       startOfDay.Weekday() == time.Saturday || startOfDay.Weekday() == time.Sunday,
			IsClosed:        isClosed,
			HasAvailability: hasAvailability,
			Slots:           slots,
		})

		currentDay = currentDay.AddDate(0, 0, 1)
	}

	return days
}

func buildSlotMeta() []slotMeta {
	meta := make([]slotMeta, 0, len(bookingSlotDefinitions))
	for _, slot := range bookingSlotDefinitions {
		meta = append(meta, slotMeta{
			ID:          slot.ID,
			Label:       slot.Label,
			Description: slot.Description,
			Duration:    fmt.Sprintf("%d hrs", int(slot.Duration.Hours())),
		})
	}
	return meta
}

func slotKey(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}

func matchSlotDefinition(start time.Time) (slotDefinition, bool) {
	startLocal := start.In(bookingLocation)
	for _, slot := range bookingSlotDefinitions {
		if slot.StartHour == startLocal.Hour() && slot.StartMinute == startLocal.Minute() {
			return slot, true
		}
	}
	return slotDefinition{}, false
}

func slotWindowLabel(start time.Time, duration time.Duration) string {
	startLocal := start.In(bookingLocation)
	endLocal := startLocal.Add(duration)
	return fmt.Sprintf("%s – %s", startLocal.Format("3:04 PM"), endLocal.Format("3:04 PM"))
}

func resolveSlotDetails(start time.Time, end time.Time) (string, string) {
	if slot, ok := matchSlotDefinition(start); ok {
		return slot.Label, slotWindowLabel(start, slot.Duration)
	}
	startLocal := start.In(bookingLocation)
	endLocal := end.In(bookingLocation)
	return "Custom Session", fmt.Sprintf("%s – %s", startLocal.Format("3:04 PM"), endLocal.Format("3:04 PM"))
}
