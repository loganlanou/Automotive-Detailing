package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "./data/detailing.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Adding missing columns to vehicles and jobs tables...")

	// Add missing columns
	migrations := []string{
		"ALTER TABLE vehicles ADD COLUMN color TEXT",
		"ALTER TABLE vehicles ADD COLUMN dealership_name TEXT",
		"ALTER TABLE vehicles ADD COLUMN dealership_logo_url TEXT",
		"ALTER TABLE vehicles ADD COLUMN dealership_location TEXT",
		"ALTER TABLE jobs ADD COLUMN slug TEXT",
		"ALTER TABLE jobs ADD COLUMN customer_testimonial TEXT",
		"ALTER TABLE jobs ADD COLUMN customer_name TEXT",
		"ALTER TABLE jobs ADD COLUMN highlight_text TEXT",
		"ALTER TABLE jobs ADD COLUMN duration_actual INTEGER",
		"ALTER TABLE jobs ADD COLUMN meta_description TEXT",
		"ALTER TABLE jobs ADD COLUMN meta_keywords TEXT",
	}

	for _, migration := range migrations {
		_, err := db.Exec(migration)
		if err != nil {
			// Ignore error if column already exists
			fmt.Printf("  Skipping: %s (may already exist)\n", migration)
		} else {
			fmt.Printf("  ✅ %s\n", migration)
		}
	}

	fmt.Println("\n✅ Migration complete!")
}
