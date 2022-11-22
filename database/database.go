package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	migratePostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	// Database connection string
	dbConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_SSL_MODE"))

	db, err := sql.Open("postgres", dbConnectionString)

	if err != nil {
		log.Fatalf("Failed to connect to database: %s\n", err)
	}

	return db
}

func MigrateConnect() *migrate.Migrate {
	db := Connect()

	driver, err := migratePostgres.WithInstance(db, &migratePostgres.Config{})

	if err != nil {
		log.Fatalf("Failed to create postgres driver: %s\n", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://../../database/migrations", "postgres", driver)

	if err != nil {
		log.Fatalf("Failed to create migrate instance: %s\n", err)
	}

	return m
}
