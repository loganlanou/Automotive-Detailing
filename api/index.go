package handler

import (
	"database/sql"
	"net/http"
	"os"

	"detailingpass/pkg/auth"
	"detailingpass/pkg/server"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "modernc.org/sqlite"
)

// Schema for C Auto Detailing Studio
var schema string = `
PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS packages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    slug TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    short_desc TEXT,
    long_desc TEXT,
    price_min INTEGER,
    price_max INTEGER,
    duration_est INTEGER,
    is_active BOOLEAN DEFAULT 1,
    sort_order INTEGER DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS gallery_groups (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    slug TEXT UNIQUE NOT NULL,
    vehicle_make TEXT,
    vehicle_model TEXT,
    vehicle_year INTEGER,
    description TEXT,
    is_featured BOOLEAN DEFAULT 0,
    sort_order INTEGER DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS media (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    gallery_group_id INTEGER,
    url TEXT NOT NULL,
    kind TEXT DEFAULT 'gallery',
    sort_order INTEGER DEFAULT 0,
    alt_text TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (gallery_group_id) REFERENCES gallery_groups(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS reviews (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    author TEXT NOT NULL,
    rating INTEGER DEFAULT 5,
    body TEXT,
    source TEXT,
    is_featured BOOLEAN DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS bookings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    customer_name TEXT NOT NULL,
    email TEXT NOT NULL,
    phone TEXT,
    vehicle_details TEXT,
    service_interest TEXT,
    notes TEXT,
    requested_start DATETIME NOT NULL,
    requested_end DATETIME NOT NULL,
    status TEXT DEFAULT 'pending',
    source TEXT DEFAULT 'web',
    internal_notes TEXT,
    clerk_user_id TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_gallery_groups_featured ON gallery_groups(is_featured);
CREATE INDEX IF NOT EXISTS idx_media_gallery_group_id ON media(gallery_group_id);
CREATE INDEX IF NOT EXISTS idx_bookings_requested_start ON bookings(requested_start);
CREATE INDEX IF NOT EXISTS idx_bookings_status ON bookings(status);
`

func Handler(w http.ResponseWriter, r *http.Request) {
	// Initialize Clerk SDK
	clerkSecretKey := os.Getenv("CLERK_SECRET_KEY")
	if clerkSecretKey != "" {
		auth.Init(clerkSecretKey)
	}

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
