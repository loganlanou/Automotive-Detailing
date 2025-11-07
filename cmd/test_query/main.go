package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"detailingpass/internal/db"

	_ "modernc.org/sqlite"
)

func main() {
	database, err := sql.Open("sqlite", "./data/detailing.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	queries := db.New(database)
	ctx := context.Background()

	// Test the query with different slugs
	slugs := []string{
		"2023-ford-f150-xlt-detail",
		"2021-subaru-outback-wilderness-detail",
	}

	for _, slug := range slugs {
		fmt.Printf("\n=== Testing slug: %s ===\n", slug)
		work, err := queries.GetWorkBySlug(ctx, sql.NullString{String: slug, Valid: true})
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
		} else {
			fmt.Printf("SUCCESS! Found job ID: %d\n", work.ID)
			fmt.Printf("  Vehicle ID: %v (valid: %v)\n", work.VehicleID.Int64, work.VehicleID.Valid)
			fmt.Printf("  Vehicle Year: %v (valid: %v)\n", work.VehicleYear.Int64, work.VehicleYear.Valid)
			fmt.Printf("  Vehicle Make: %s (valid: %v)\n", work.VehicleMake.String, work.VehicleMake.Valid)
			fmt.Printf("  Vehicle Model: %s (valid: %v)\n", work.VehicleModel.String, work.VehicleModel.Valid)
			fmt.Printf("  Vehicle Trim: %s (valid: %v)\n", work.VehicleTrim.String, work.VehicleTrim.Valid)
			fmt.Printf("  Vehicle Color: %s (valid: %v)\n", work.VehicleColor.String, work.VehicleColor.Valid)
			fmt.Printf("  Dealership Name: %s (valid: %v)\n", work.DealershipName.String, work.DealershipName.Valid)
			fmt.Printf("  Dealership Location: %s (valid: %v)\n", work.DealershipLocation.String, work.DealershipLocation.Valid)
			fmt.Printf("  Package Name: %s (valid: %v)\n", work.PackageName.String, work.PackageName.Valid)
		}
	}

	// Also list all jobs to see what slugs exist
	fmt.Printf("\n=== All jobs in database ===\n")
	rows, err := database.Query("SELECT id, slug FROM jobs")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var slug sql.NullString
		rows.Scan(&id, &slug)
		fmt.Printf("Job %d: %s\n", id, slug.String)
	}
}
