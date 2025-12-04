package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"detailingpass/pkg/db"
	"detailingpass/web/templates/pages"

	"github.com/labstack/echo/v4"
)

// GalleryItem represents a gallery group with its images for the template
type GalleryItem struct {
	ID          int64
	Slug        string
	Title       string
	VehicleMake string
	VehicleModel string
	VehicleYear int64
	Description string
	IsFeatured  bool
	HeroImage   string
	Images      []GalleryImage
}

type GalleryImage struct {
	ID       int64
	URL      string
	Kind     string
	AltText  string
}

func (h *Handler) Gallery(c echo.Context) error {
	ctx := c.Request().Context()
	queries := db.New(h.db)

	// Get all gallery groups
	groups, err := queries.ListGalleryGroups(ctx, db.ListGalleryGroupsParams{
		Limit:  100,
		Offset: 0,
	})
	if err != nil {
		c.Logger().Errorf("Failed to fetch gallery groups: %v", err)
		groups = []db.GalleryGroup{}
	}

	// Build gallery items with images
	var items []pages.GalleryItem
	for _, g := range groups {
		media, mediaErr := queries.GetMediaForGalleryGroup(ctx, sql.NullInt64{Int64: g.ID, Valid: true})
		if mediaErr != nil {
			c.Logger().Warnf("Failed to fetch media for gallery group %d: %v", g.ID, mediaErr)
		}

		var images []pages.GalleryImage
		heroImage := "/static/images/placeholder.jpg"

		for _, m := range media {
			img := pages.GalleryImage{
				ID:      m.ID,
				URL:     m.Url,
				Kind:    m.Kind.String,
				AltText: m.AltText.String,
			}
			images = append(images, img)
			if m.Kind.String == "hero" && heroImage == "/static/images/placeholder.jpg" {
				heroImage = m.Url
			}
		}

		// Use first image as hero if no hero designated
		if heroImage == "/static/images/placeholder.jpg" && len(images) > 0 {
			heroImage = images[0].URL
		}

		items = append(items, pages.GalleryItem{
			ID:           g.ID,
			Slug:         g.Slug,
			Title:        g.Title,
			VehicleMake:  g.VehicleMake.String,
			VehicleModel: g.VehicleModel.String,
			VehicleYear:  g.VehicleYear.Int64,
			Description:  g.Description.String,
			IsFeatured:   g.IsFeatured.Bool,
			HeroImage:    heroImage,
			Images:       images,
		})
	}

	return pages.Gallery(items).Render(c.Request().Context(), c.Response().Writer)
}

// Admin Gallery handlers

func (h *Handler) AdminGallery(c echo.Context) error {
	ctx := c.Request().Context()
	queries := db.New(h.db)

	groups, err := queries.ListGalleryGroups(ctx, db.ListGalleryGroupsParams{
		Limit:  100,
		Offset: 0,
	})
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to fetch gallery groups")
	}

	// Check if we're editing
	var formData *pages.GalleryFormData
	editID := c.QueryParam("edit")
	if editID != "" {
		id, err := strconv.ParseInt(editID, 10, 64)
		if err == nil {
			group, err := queries.GetGalleryGroupByID(ctx, id)
			if err == nil {
				formData = &pages.GalleryFormData{
					ID:           group.ID,
					Title:        group.Title,
					Slug:         group.Slug,
					VehicleMake:  group.VehicleMake.String,
					VehicleModel: group.VehicleModel.String,
					VehicleYear:  group.VehicleYear.Int64,
					Description:  group.Description.String,
					IsFeatured:   group.IsFeatured.Bool,
					SortOrder:    group.SortOrder.Int64,
					IsEdit:       true,
				}
			}
		}
	}

	return pages.AdminGallery(groups, formData).Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) CreateGalleryGroup(c echo.Context) error {
	ctx := c.Request().Context()
	queries := db.New(h.db)

	title := c.FormValue("title")
	slug := c.FormValue("slug")
	if slug == "" {
		slug = strings.ToLower(strings.ReplaceAll(title, " ", "-"))
	}
	vehicleMake := c.FormValue("vehicle_make")
	vehicleModel := c.FormValue("vehicle_model")
	vehicleYearStr := c.FormValue("vehicle_year")
	vehicleYear, _ := strconv.ParseInt(vehicleYearStr, 10, 64)
	description := c.FormValue("description")
	isFeatured := c.FormValue("is_featured") == "true"
	sortOrderStr := c.FormValue("sort_order")
	sortOrder, _ := strconv.ParseInt(sortOrderStr, 10, 64)

	_, err := queries.CreateGalleryGroup(ctx, db.CreateGalleryGroupParams{
		Title:        title,
		Slug:         slug,
		VehicleMake:  sql.NullString{String: vehicleMake, Valid: vehicleMake != ""},
		VehicleModel: sql.NullString{String: vehicleModel, Valid: vehicleModel != ""},
		VehicleYear:  sql.NullInt64{Int64: vehicleYear, Valid: vehicleYear > 0},
		Description:  sql.NullString{String: description, Valid: description != ""},
		IsFeatured:   sql.NullBool{Bool: isFeatured, Valid: true},
		SortOrder:    sql.NullInt64{Int64: sortOrder, Valid: true},
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to create gallery group: %v", err))
	}

	return c.Redirect(http.StatusSeeOther, "/admin/gallery")
}

func (h *Handler) UpdateGalleryGroup(c echo.Context) error {
	ctx := c.Request().Context()
	queries := db.New(h.db)

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid gallery group ID")
	}

	title := c.FormValue("title")
	slug := c.FormValue("slug")
	vehicleMake := c.FormValue("vehicle_make")
	vehicleModel := c.FormValue("vehicle_model")
	vehicleYearStr := c.FormValue("vehicle_year")
	vehicleYear, _ := strconv.ParseInt(vehicleYearStr, 10, 64)
	description := c.FormValue("description")
	isFeatured := c.FormValue("is_featured") == "true"
	sortOrderStr := c.FormValue("sort_order")
	sortOrder, _ := strconv.ParseInt(sortOrderStr, 10, 64)

	_, err = queries.UpdateGalleryGroup(ctx, db.UpdateGalleryGroupParams{
		ID:           id,
		Title:        title,
		Slug:         slug,
		VehicleMake:  sql.NullString{String: vehicleMake, Valid: vehicleMake != ""},
		VehicleModel: sql.NullString{String: vehicleModel, Valid: vehicleModel != ""},
		VehicleYear:  sql.NullInt64{Int64: vehicleYear, Valid: vehicleYear > 0},
		Description:  sql.NullString{String: description, Valid: description != ""},
		IsFeatured:   sql.NullBool{Bool: isFeatured, Valid: true},
		SortOrder:    sql.NullInt64{Int64: sortOrder, Valid: true},
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to update gallery group: %v", err))
	}

	return c.Redirect(http.StatusSeeOther, "/admin/gallery")
}

func (h *Handler) DeleteGalleryGroup(c echo.Context) error {
	ctx := c.Request().Context()
	queries := db.New(h.db)

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid gallery group ID")
	}

	err = queries.DeleteGalleryGroup(ctx, id)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to delete gallery group: %v", err))
	}

	return c.Redirect(http.StatusSeeOther, "/admin/gallery")
}
