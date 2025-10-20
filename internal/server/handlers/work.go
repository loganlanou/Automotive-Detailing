package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"detailingpass/internal/db"
	"detailingpass/web/templates/pages"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Work(c echo.Context) error {
	ctx := context.Background()
	queries := db.New(h.db)

	// TODO: Parse filter params (make, model, year, package, price_min, price_max)
	// For now, just fetch all work with pagination
	limit := int64(12)
	offset := int64(0)

	// Fetch all work items
	workRows, err := queries.ListAllWork(ctx, db.ListAllWorkParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to fetch work")
	}

	// Convert to WorkListItem format and fetch primary images
	workItems := make([]pages.WorkListItem, 0, len(workRows))
	for _, row := range workRows {
		// Get the first "after" or "hero" image for this job
		var primaryImageURL string
		if row.ID > 0 {
			media, err := queries.GetMediaForJob(ctx, sql.NullInt64{Int64: row.ID, Valid: true})
			if err == nil && len(media) > 0 {
				// Find first "after" or "hero" image
				for _, m := range media {
					if m.Kind.String == "after" || m.Kind.String == "hero" {
						primaryImageURL = m.Url
						break
					}
				}
				// Fallback to first image
				if primaryImageURL == "" && len(media) > 0 {
					primaryImageURL = media[0].Url
				}
			}
		}

		workItems = append(workItems, pages.WorkListItem{
			ID:              row.ID,
			Slug:            row.Slug,
			Featured:        row.Featured,
			HighlightText:   row.HighlightText,
			DisplayPrice:    row.DisplayPrice,
			CompletedAt:     row.CompletedAt,
			VehicleYear:     row.VehicleYear,
			VehicleMake:     row.VehicleMake,
			VehicleModel:    row.VehicleModel,
			VehicleTrim:     row.VehicleTrim,
			DealershipName:  row.DealershipName,
			PackageName:     row.PackageName,
			PrimaryImageURL: primaryImageURL,
		})
	}

	return pages.Work(workItems).Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) WorkDetail(c echo.Context) error {
	slug := c.Param("slug")
	ctx := context.Background()
	queries := db.New(h.db)

	// Fetch work detail by slug
	// Note: This query doesn't exist yet in our queries, so we'll need to add it or use GetJobByID
	// For now, let's return a placeholder
	_ = slug

	// TODO: Implement GetWorkBySlug query and use it here
	// This is just a placeholder structure for now
	data := pages.WorkDetailData{
		Job:            db.Job{},
		VehicleYear:    sql.NullInt64{},
		VehicleMake:    sql.NullString{},
		VehicleModel:   sql.NullString{},
		VehicleTrim:    sql.NullString{},
		VehicleColor:   sql.NullString{},
		DealershipName: sql.NullString{},
		BeforeImages:   []db.Medium{},
		AfterImages:    []db.Medium{},
		GalleryImages:  []db.Medium{},
		RelatedWork:    []pages.WorkListItem{},
	}

	// Fetch all jobs to find one by slug (temporary workaround)
	allJobs, err := queries.ListFeaturedJobs(ctx, 100)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to fetch work details")
	}

	var foundJobID int64
	var foundVehicleID sql.NullInt64
	var foundPackageID sql.NullInt64
	foundJobExists := false

	for _, job := range allJobs {
		if job.Slug.Valid && job.Slug.String == slug {
			foundJobID = job.ID
			foundVehicleID = job.VehicleID
			foundPackageID = job.PackageID
			foundJobExists = true
			break
		}
	}

	if !foundJobExists {
		return c.String(http.StatusNotFound, "Work not found")
	}

	// Get full job details
	foundJob, err := queries.GetJobByID(ctx, foundJobID)
	if err != nil {
		return c.String(http.StatusNotFound, "Work not found")
	}

	data.Job = foundJob

	// Fetch vehicle details
	if foundVehicleID.Valid {
		vehicle, err := queries.GetVehicleBySlug(ctx, fmt.Sprintf("vehicle-%d", foundVehicleID.Int64))
		if err == nil {
			data.VehicleYear = vehicle.Year
			data.VehicleMake = sql.NullString{String: vehicle.Make, Valid: true}
			data.VehicleModel = sql.NullString{String: vehicle.Model, Valid: true}
			data.VehicleTrim = sql.NullString{String: vehicle.Trim.String, Valid: vehicle.Trim.Valid}
			data.VehicleColor = vehicle.Color
			data.DealershipName = vehicle.DealershipName
			data.DealershipLogoURL = vehicle.DealershipLogoUrl
			data.DealershipListingURL = vehicle.DealershipListingUrl
			data.DealershipLocation = vehicle.DealershipLocation
		}
	}

	// Fetch package details
	if foundPackageID.Valid {
		pkg, err := queries.GetPackageByID(ctx, foundPackageID.Int64)
		if err == nil {
			data.PackageName = sql.NullString{String: pkg.Name, Valid: true}
			data.PackageShortDesc = pkg.ShortDesc
			data.PackagePriceMin = pkg.PriceMin
			data.PackagePriceMax = pkg.PriceMax
		}
	}

	// Fetch media for this job
	media, err := queries.GetMediaForJob(ctx, sql.NullInt64{Int64: foundJob.ID, Valid: true})
	if err == nil {
		for _, m := range media {
			switch m.Kind.String {
			case "before":
				data.BeforeImages = append(data.BeforeImages, m)
			case "after", "hero":
				data.AfterImages = append(data.AfterImages, m)
			case "gallery":
				data.GalleryImages = append(data.GalleryImages, m)
			default:
				data.GalleryImages = append(data.GalleryImages, m)
			}
		}
	}

	// Fetch related work (same make or same package)
	if foundVehicleID.Valid {
		vehicle, _ := queries.GetVehicleBySlug(ctx, fmt.Sprintf("vehicle-%d", foundVehicleID.Int64))
		relatedRows, err := queries.GetRelatedWork(ctx, db.GetRelatedWorkParams{
			ID:        foundJobID,
			Make:      vehicle.Make,
			PackageID: foundPackageID,
		})
		if err == nil {
			for _, row := range relatedRows {
				// Get primary image
				var primaryImageURL string
				if row.ID > 0 {
					relatedMedia, err := queries.GetMediaForJob(ctx, sql.NullInt64{Int64: row.ID, Valid: true})
					if err == nil && len(relatedMedia) > 0 {
						for _, m := range relatedMedia {
							if m.Kind.String == "after" || m.Kind.String == "hero" {
								primaryImageURL = m.Url
								break
							}
						}
						if primaryImageURL == "" {
							primaryImageURL = relatedMedia[0].Url
						}
					}
				}

				data.RelatedWork = append(data.RelatedWork, pages.WorkListItem{
					ID:              row.ID,
					Slug:            row.Slug,
					Featured:        row.Featured,
					CompletedAt:     row.CompletedAt,
					VehicleYear:     row.VehicleYear,
					VehicleMake:     row.VehicleMake,
					VehicleModel:    row.VehicleModel,
					DealershipName:  row.DealershipName,
					PackageName:     row.PackageName,
					PrimaryImageURL: primaryImageURL,
				})
			}
		}
	}

	return pages.WorkDetail(data).Render(c.Request().Context(), c.Response().Writer)
}
