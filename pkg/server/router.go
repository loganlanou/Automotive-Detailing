package server

import (
	"database/sql"

	"detailingpass/pkg/server/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, db *sql.DB) {
	h := handlers.New(db)

	// Health check endpoint for Vercel uptime probes
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})

	// Public pages
	e.GET("/", h.Home)
	e.GET("/about", h.About)
	e.GET("/services", h.Services)
	e.GET("/services/:slug", h.ServiceDetail)
	e.GET("/work", h.Work)
	e.GET("/work/:slug", h.WorkDetail)
	e.GET("/reviews", h.Reviews)
	e.GET("/contact", h.Contact)
	e.POST("/contact", h.ContactSubmit)
	e.GET("/blog", h.Blog)
	e.GET("/blog/:slug", h.BlogPost)
	e.GET("/faq", h.FAQ)
	e.GET("/privacy", h.Privacy)
	e.GET("/terms", h.Terms)

	// Design system showcase
	e.GET("/style", h.StyleGuide)

	// Admin routes (to be protected with Clerk middleware)
	admin := e.Group("/admin")
	// TODO: Add Clerk auth middleware here
	admin.GET("", h.AdminDashboard)
	admin.GET("/packages", h.AdminPackages)
	admin.POST("/packages", h.CreatePackage)
	admin.POST("/packages/:id", h.UpdatePackage)
	admin.POST("/packages/:id/delete", h.DeletePackage)
	admin.GET("/vehicles", h.AdminVehicles)
	admin.GET("/jobs", h.AdminJobs)
	admin.GET("/media", h.AdminMedia)

	// API routes
	api := e.Group("/api")
	api.POST("/dealer/sync", h.DealerSync)
	api.GET("/dealer/export", h.DealerExport)
}
