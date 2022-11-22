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
		log.Fatal("[Error] Failed to load .env file")
	}

	m := database.MigrateConnect()

	if err := m.Up(); err != nil {
		if strings.Compare(fmt.Sprint(err), "no change") == 0 {
			fmt.Println("Migration executed successfully but with no change.")
			return
		} else {
			fmt.Printf("An error occured while executing the migration: %s\n", err)
			return
		}
	}

	fmt.Println("Migration executed successfully.")
}
