package server

import (
	"database/sql"
	"os"

	"detailingpass/pkg/auth"
	"detailingpass/pkg/server/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, db *sql.DB) {
	// Initialize Clerk SDK
	clerkSecretKey := os.Getenv("CLERK_SECRET_KEY")
	if clerkSecretKey != "" {
		auth.Init(clerkSecretKey)
	}

	h := handlers.New(db)

	// Health check endpoint for Vercel uptime probes
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})

	// Public pages
	e.GET("/", h.Home)
	e.GET("/gallery", h.Gallery)
	e.GET("/about", h.About)
	e.GET("/booking", h.BookingPage)
	e.GET("/privacy", h.Privacy)
	e.GET("/terms", h.Terms)

	// Auth pages
	e.GET("/sign-in", h.SignIn)
	e.GET("/sign-up", h.SignUp)

	// Admin routes (protected with Clerk middleware)
	admin := e.Group("/admin")
	admin.Use(auth.RequireAuth())
	admin.GET("", h.AdminDashboard)
	admin.GET("/packages", h.AdminPackages)
	admin.POST("/packages", h.CreatePackage)
	admin.POST("/packages/:id", h.UpdatePackage)
	admin.POST("/packages/:id/delete", h.DeletePackage)
	admin.GET("/bookings", h.AdminBookings)
	admin.POST("/bookings/:id/status", h.UpdateBookingStatus)
	admin.GET("/gallery", h.AdminGallery)
	admin.POST("/gallery", h.CreateGalleryGroup)
	admin.POST("/gallery/:id", h.UpdateGalleryGroup)
	admin.POST("/gallery/:id/delete", h.DeleteGalleryGroup)

	// API routes (with optional auth to capture user ID if logged in)
	api := e.Group("/api")
	api.Use(auth.OptionalAuth())
	api.GET("/bookings/availability", h.BookingAvailability)
	api.POST("/bookings", h.CreateBookingRequest)
}
