package handlers

import (
	"context"
	"database/sql"
	"net/http"

	"detailingpass/pkg/db"
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
		// If error, return empty list instead of error (database might be empty)
		workRows = []db.ListAllWorkRow{}
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

	// Fetch work detail by slug using the proper query
	work, err := queries.GetWorkBySlug(ctx, sql.NullString{String: slug, Valid: true})
	if err != nil {
		return c.String(http.StatusNotFound, "Work not found")
	}

	// Initialize the data structure - construct Job from GetWorkBySlugRow fields
	data := pages.WorkDetailData{
		Job: db.Job{
			ID:                  work.ID,
			Slug:                work.Slug,
			VehicleID:           work.VehicleID,
			PackageID:           work.PackageID,
			Technician:          work.Technician,
			Notes:               work.Notes,
			CompletedAt:         work.CompletedAt,
			DurationActual:      work.DurationActual,
			Featured:            work.Featured,
			DisplayPrice:        work.DisplayPrice,
			HighlightText:       work.HighlightText,
			CustomerTestimonial: work.CustomerTestimonial,
			CustomerName:        work.CustomerName,
			MetaDescription:     work.MetaDescription,
			MetaKeywords:        work.MetaKeywords,
			CreatedAt:           work.CreatedAt,
			UpdatedAt:           work.UpdatedAt,
		},
		VehicleYear:          work.VehicleYear,
		VehicleMake:          work.VehicleMake,
		VehicleModel:         work.VehicleModel,
		VehicleTrim:          work.VehicleTrim,
		VehicleColor:         work.VehicleColor,
		DealershipName:       work.DealershipName,
		DealershipLogoURL:    work.DealershipLogoUrl,
		DealershipListingURL: work.DealershipListingUrl,
		DealershipLocation:   work.DealershipLocation,
		PackageName:          work.PackageName,
		PackageShortDesc:     work.PackageShortDesc,
		PackagePriceMin:      work.PackagePriceMin,
		PackagePriceMax:      work.PackagePriceMax,
		BeforeImages:         []db.Medium{},
		AfterImages:          []db.Medium{},
		GalleryImages:        []db.Medium{},
		RelatedWork:          []pages.WorkListItem{},
	}

	// Fetch media for this job
	media, err := queries.GetMediaForJob(ctx, sql.NullInt64{Int64: work.ID, Valid: true})
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
	if work.VehicleMake.Valid && work.PackageID.Valid {
		relatedRows, err := queries.GetRelatedWork(ctx, db.GetRelatedWorkParams{
			ID:        work.ID,
			Make:      work.VehicleMake.String,
			PackageID: work.PackageID,
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
