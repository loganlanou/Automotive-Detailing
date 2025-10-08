package handler

import (
	"database/sql"
	"net/http"

	"detailingpass/internal/server"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	_ "embed"
)

//go:embed ../cmd/server/schema.sql
var schema string

func Handler(w http.ResponseWriter, r *http.Request) {
	// Create Echo instance
	e := echo.New()
	e.HideBanner = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Static files
	e.Static("/static", "web/static")

	// Initialize database
	db, _ := sql.Open("sqlite3", ":memory:")
	db.Exec(schema)

	// Setup routes
	server.SetupRoutes(e, db)

	// Serve the request
	e.ServeHTTP(w, r)
}
