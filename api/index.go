package handler

import (
	"database/sql"
	"net/http"

	"detailingpass/internal/server"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "modernc.org/sqlite"
	_ "embed"
)

// Schema is defined inline to avoid embed path issues
var schema string = `
CREATE TABLE IF NOT EXISTS contacts (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT NOT NULL,
	phone TEXT,
	service TEXT,
	message TEXT NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
`

func Handler(w http.ResponseWriter, r *http.Request) {
	// Create Echo instance
	e := echo.New()
	e.HideBanner = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Static files
	e.Static("/static", "web/static")
	e.Static("/uploads", "web/static/uploads")
	e.File("/favicon.png", "public/favicon.png")
	e.File("/favicon.ico", "public/favicon.png")
	e.File("/robots.txt", "public/robots.txt")
	e.File("/sitemap.xml", "public/sitemap.xml")

	// Initialize database
	db, _ := sql.Open("sqlite", ":memory:")
	db.Exec(schema)

	// Setup routes
	server.SetupRoutes(e, db)

	// Serve the request
	e.ServeHTTP(w, r)
}
