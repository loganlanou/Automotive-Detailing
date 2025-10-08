package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) AdminDashboard(c echo.Context) error {
	// TODO: Render admin dashboard template
	return c.String(http.StatusOK, "Admin Dashboard - TODO")
}

func (h *Handler) AdminPackages(c echo.Context) error {
	// TODO: CRUD for packages
	return c.String(http.StatusOK, "Admin Packages - TODO")
}

func (h *Handler) AdminVehicles(c echo.Context) error {
	// TODO: CRUD for vehicles
	return c.String(http.StatusOK, "Admin Vehicles - TODO")
}

func (h *Handler) AdminJobs(c echo.Context) error {
	// TODO: CRUD for jobs
	return c.String(http.StatusOK, "Admin Jobs - TODO")
}

func (h *Handler) AdminMedia(c echo.Context) error {
	// TODO: Upload and manage media
	return c.String(http.StatusOK, "Admin Media - TODO")
}

func (h *Handler) DealerSync(c echo.Context) error {
	// TODO: Accept JSON payload for vehicle + job sync
	// TODO: Create or update vehicle and job records
	// TODO: Log sync attempt
	return c.JSON(http.StatusOK, map[string]string{
		"status": "synced",
	})
}

func (h *Handler) DealerExport(c echo.Context) error {
	// TODO: Export recent jobs as CSV
	c.Response().Header().Set("Content-Type", "text/csv")
	c.Response().Header().Set("Content-Disposition", "attachment; filename=dealer-export.csv")

	csv := "VIN,Make,Model,Year,Package,Completed,URL\n"
	// TODO: Add actual data

	return c.String(http.StatusOK, csv)
}
