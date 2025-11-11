package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"detailingpass/pkg/db"
	"detailingpass/web/templates/pages"

	"github.com/labstack/echo/v4"
)

func (h *Handler) AdminDashboard(c echo.Context) error {
	ctx := context.Background()
	queries := db.New(h.db)

	// Get counts for all entities
	packageCount, _ := queries.CountPackages(ctx)
	vehicleCount, _ := queries.CountVehicles(ctx)
	jobCount, _ := queries.CountJobs(ctx)
	mediaCount, _ := queries.CountMedia(ctx)
	reviewCount, _ := queries.CountReviews(ctx)
	postCount, _ := queries.CountPosts(ctx)

	stats := pages.DashboardStats{
		PackageCount: packageCount,
		VehicleCount: vehicleCount,
		JobCount:     jobCount,
		MediaCount:   mediaCount,
		ReviewCount:  reviewCount,
		PostCount:    postCount,
	}

	return pages.AdminDashboard(stats).Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) AdminPackages(c echo.Context) error {
	ctx := context.Background()
	queries := db.New(h.db)

	// Get all packages
	packages, err := queries.GetAllPackagesAdmin(ctx)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to fetch packages")
	}

	// Check if we're editing a package
	var formData *pages.PackageFormData
	editID := c.QueryParam("edit")
	if editID != "" {
		id, err := strconv.ParseInt(editID, 10, 64)
		if err == nil {
			pkg, err := queries.GetPackageByID(ctx, id)
			if err == nil {
				formData = &pages.PackageFormData{
					ID:          pkg.ID,
					Slug:        pkg.Slug,
					Name:        pkg.Name,
					ShortDesc:   pkg.ShortDesc.String,
					LongDesc:    pkg.LongDesc.String,
					PriceMin:    pkg.PriceMin.Int64,
					PriceMax:    pkg.PriceMax.Int64,
					DurationEst: pkg.DurationEst.Int64,
					IsActive:    pkg.IsActive.Bool,
					SortOrder:   pkg.SortOrder.Int64,
					IsEdit:      true,
				}
			}
		}
	}

	return pages.AdminPackages(packages, formData).Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) CreatePackage(c echo.Context) error {
	ctx := context.Background()
	queries := db.New(h.db)

	// Parse form data
	name := c.FormValue("name")
	slug := c.FormValue("slug")
	shortDesc := c.FormValue("short_desc")
	longDesc := c.FormValue("long_desc")

	// Parse prices (convert from dollars to cents)
	priceMinStr := c.FormValue("price_min")
	priceMaxStr := c.FormValue("price_max")
	priceMin, _ := strconv.ParseFloat(priceMinStr, 64)
	priceMax, _ := strconv.ParseFloat(priceMaxStr, 64)
	priceMinCents := int64(priceMin * 100)
	priceMaxCents := int64(priceMax * 100)

	// Parse duration (convert from hours to minutes)
	durationStr := c.FormValue("duration_est")
	durationHours, _ := strconv.ParseFloat(durationStr, 64)
	durationMinutes := int64(durationHours * 60)

	// Parse sort order
	sortOrderStr := c.FormValue("sort_order")
	sortOrder, _ := strconv.ParseInt(sortOrderStr, 10, 64)

	// Parse is_active checkbox
	isActive := c.FormValue("is_active") == "true"

	// Create package
	_, err := queries.CreatePackage(ctx, db.CreatePackageParams{
		Slug:        slug,
		Name:        name,
		ShortDesc:   sql.NullString{String: shortDesc, Valid: shortDesc != ""},
		LongDesc:    sql.NullString{String: longDesc, Valid: longDesc != ""},
		PriceMin:    sql.NullInt64{Int64: priceMinCents, Valid: true},
		PriceMax:    sql.NullInt64{Int64: priceMaxCents, Valid: true},
		DurationEst: sql.NullInt64{Int64: durationMinutes, Valid: true},
		IsActive:    sql.NullBool{Bool: isActive, Valid: true},
		SortOrder:   sql.NullInt64{Int64: sortOrder, Valid: true},
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to create package: %v", err))
	}

	return c.Redirect(http.StatusSeeOther, "/admin/packages")
}

func (h *Handler) UpdatePackage(c echo.Context) error {
	ctx := context.Background()
	queries := db.New(h.db)

	// Get package ID from URL
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid package ID")
	}

	// Parse form data
	name := c.FormValue("name")
	slug := c.FormValue("slug")
	shortDesc := c.FormValue("short_desc")
	longDesc := c.FormValue("long_desc")

	// Parse prices (convert from dollars to cents)
	priceMinStr := c.FormValue("price_min")
	priceMaxStr := c.FormValue("price_max")
	priceMin, _ := strconv.ParseFloat(priceMinStr, 64)
	priceMax, _ := strconv.ParseFloat(priceMaxStr, 64)
	priceMinCents := int64(priceMin * 100)
	priceMaxCents := int64(priceMax * 100)

	// Parse duration (convert from hours to minutes)
	durationStr := c.FormValue("duration_est")
	durationHours, _ := strconv.ParseFloat(durationStr, 64)
	durationMinutes := int64(durationHours * 60)

	// Parse sort order
	sortOrderStr := c.FormValue("sort_order")
	sortOrder, _ := strconv.ParseInt(sortOrderStr, 10, 64)

	// Parse is_active checkbox
	isActive := c.FormValue("is_active") == "true"

	// Update package
	_, err = queries.UpdatePackage(ctx, db.UpdatePackageParams{
		ID:          id,
		Slug:        slug,
		Name:        name,
		ShortDesc:   sql.NullString{String: shortDesc, Valid: shortDesc != ""},
		LongDesc:    sql.NullString{String: longDesc, Valid: longDesc != ""},
		PriceMin:    sql.NullInt64{Int64: priceMinCents, Valid: true},
		PriceMax:    sql.NullInt64{Int64: priceMaxCents, Valid: true},
		DurationEst: sql.NullInt64{Int64: durationMinutes, Valid: true},
		IsActive:    sql.NullBool{Bool: isActive, Valid: true},
		SortOrder:   sql.NullInt64{Int64: sortOrder, Valid: true},
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to update package: %v", err))
	}

	return c.Redirect(http.StatusSeeOther, "/admin/packages")
}

func (h *Handler) DeletePackage(c echo.Context) error {
	ctx := context.Background()
	queries := db.New(h.db)

	// Get package ID from URL
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid package ID")
	}

	// Delete package
	err = queries.DeletePackage(ctx, id)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to delete package: %v", err))
	}

	return c.Redirect(http.StatusSeeOther, "/admin/packages")
}

func (h *Handler) AdminVehicles(c echo.Context) error {
	// TODO: CRUD for vehicles
	return c.String(http.StatusOK, "Admin Vehicles - Coming Soon")
}

func (h *Handler) AdminJobs(c echo.Context) error {
	// TODO: CRUD for jobs
	return c.String(http.StatusOK, "Admin Jobs - Coming Soon")
}

func (h *Handler) AdminMedia(c echo.Context) error {
	// TODO: Upload and manage media
	return c.String(http.StatusOK, "Admin Media - Coming Soon")
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
