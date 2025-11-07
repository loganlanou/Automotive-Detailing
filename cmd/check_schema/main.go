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

	// Get table info for vehicles
	rows, err := db.Query("PRAGMA table_info(vehicles)")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Vehicles table columns:")
	for rows.Next() {
		var cid int
		var name, typ string
		var notnull, dfltValue, pk interface{}
		err = rows.Scan(&cid, &name, &typ, &notnull, &dfltValue, &pk)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("  - %s (%s)\n", name, typ)
	}
}
