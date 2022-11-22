package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chrislentz/example-go-restful-api/database"
	"github.com/joho/godotenv"
)

func main() {
	// Enable line numbers in logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Load .env variables
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("[Error] Failed to load .env file")
	}

	if os.Getenv("ENVIRONMENT") != "production" {
		db := database.Connect()

		if _, err := db.Exec("DO $$ DECLARE r RECORD;BEGIN FOR r IN(SELECT tablename FROM pg_tables WHERE schemaname=current_schema())LOOP EXECUTE'DROP TABLE '||quote_ident(r.tablename)||' CASCADE';END LOOP;END $$;"); err != nil {
			log.Fatalf("[ERROR] %s", err)
			return
		}

		fmt.Println("Drop executed successfully.")
	} else {
		fmt.Println("Drop not allowed on production.")
	}
}
