package main

import (
	"fmt"
	"log"

	"github.com/chrislentz/example-go-restful-api/database"
	"github.com/chrislentz/example-go-restful-api/database/seeders"
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

	db := database.Connect()

	if !seeders.UsersSeeder(db) {
		fmt.Println("Roles seeder executed successfully.")
	}
}
