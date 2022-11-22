package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/chrislentz/example-go-restful-api/database"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Enable line numbers in logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Load .env variables
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	m := database.MigrateConnect()

	if err := m.Steps(-1); err != nil {
		if strings.Compare(fmt.Sprint(err), "file does not exist") == 0 {
			fmt.Println("Rollback executed successfully but with no change.")
			return
		} else {
			log.Fatalf("An error occured while executing the migration: %s\n", err)
		}
	}

	version, _, _ := m.Version()

	fmt.Printf("Rollback executed successfully. Current Migration Version: %d\n", version)
}
