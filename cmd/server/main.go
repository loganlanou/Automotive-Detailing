package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"detailingpass/pkg/server"

	_ "embed"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "modernc.org/sqlite"
)

//go:embed schema.sql
var schema string

func runMigrations(db *sql.DB) error {
	_, err := db.Exec(schema)
	return err
}

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Get port from env or default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Database setup
	dbPath := os.Getenv("DATABASE_PATH")
	if dbPath == "" {
		dbPath = "./data/detailing.db"
	}

	// Ensure data directory exists
	if err := os.MkdirAll("./data", 0755); err != nil {
		log.Fatalf("Failed to create data directory: %v", err)
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations
	if err := runMigrations(db); err != nil {
		log.Printf("Warning: Failed to run migrations: %v", err)
	}

	// Echo setup
	e := echo.New()
	e.HideBanner = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.CORS())

	// Static files
	e.Static("/static", "web/static")
	e.Static("/uploads", "web/static/uploads")
	e.File("/favicon.png", "public/favicon.png")
	e.File("/favicon.ico", "public/favicon.png")
	e.File("/robots.txt", "public/robots.txt")
	e.File("/sitemap.xml", "public/sitemap.xml")

	// Setup routes
	server.SetupRoutes(e, db)

	// Start server
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	log.Printf("🚀 Server starting on %s", addr)
	if err := e.Start(addr); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
