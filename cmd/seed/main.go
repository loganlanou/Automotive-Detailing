package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "modernc.org/sqlite"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Open database
	db, err := sql.Open("sqlite", "./data/detailing.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Clear existing data first
	fmt.Println("Clearing existing vehicles and jobs...")
	_, err = db.Exec("DELETE FROM media")
	if err != nil {
		log.Fatal("Error clearing media:", err)
	}
	_, err = db.Exec("DELETE FROM jobs")
	if err != nil {
		log.Fatal("Error clearing jobs:", err)
	}
	_, err = db.Exec("DELETE FROM vehicles")
	if err != nil {
		log.Fatal("Error clearing vehicles:", err)
	}
	_, err = db.Exec("DELETE FROM packages")
	if err != nil {
		log.Fatal("Error clearing packages:", err)
	}

	// Reset auto-increment counters
	_, err = db.Exec("DELETE FROM sqlite_sequence WHERE name IN ('vehicles', 'jobs', 'media', 'packages')")
	if err != nil {
		log.Fatal("Error resetting auto-increment:", err)
	}

	// Read seed file
	seedSQL, err := os.ReadFile("./pkg/db/seed_vehicles.sql")
	if err != nil {
		log.Fatal(err)
	}

	// Execute seed SQL with multiple statements support
	fmt.Println("Inserting seed data...")

	// Split into statements (simple split by semicolon)
	statements := []string{}
	currentStmt := ""
	for _, line := range strings.Split(string(seedSQL), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "--") {
			continue
		}
		currentStmt += " " + line
		if strings.HasSuffix(line, ";") {
			statements = append(statements, currentStmt)
			currentStmt = ""
		}
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	for i, stmt := range statements {
		if strings.TrimSpace(stmt) == "" {
			continue
		}
		_, err = tx.Exec(stmt)
		if err != nil {
			tx.Rollback()
			log.Fatalf("Error executing statement %d: %v\nStatement: %s", i+1, err, stmt[:min(len(stmt), 200)])
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("Error committing transaction:", err)
	}

	// Verify data was inserted
	var vehicleCount, jobCount int
	err = db.QueryRow("SELECT COUNT(*) FROM vehicles").Scan(&vehicleCount)
	if err != nil {
		log.Fatal(err)
	}
	err = db.QueryRow("SELECT COUNT(*) FROM jobs").Scan(&jobCount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("✅ Seed data inserted successfully!\n")
	fmt.Printf("   - Vehicles: %d\n", vehicleCount)
	fmt.Printf("   - Jobs: %d\n", jobCount)
	fmt.Println("\nYour gallery should now show 12 vehicles:")
	fmt.Println("   - 6 from Courtesy Automotive Group")
	fmt.Println("   - 6 customer vehicles")
}
