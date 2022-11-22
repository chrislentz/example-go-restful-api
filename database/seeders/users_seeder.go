package seeders

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/chrislentz/example-go-restful-api/sqlc"
)

func UsersSeeder(db *sql.DB) bool {
	var err error
	ctx := context.Background()
	queries := sqlc.New(db)

	// Only execute seeder if users table is empty
	users_count, err := queries.CountUsers(ctx)
	if err != nil {
		log.Fatalf("[ERROR] %s", err)
	}

	if users_count != 0 {
		return false
	}

	if os.Getenv("ENVIRONMENT") == "local" {
		// Start Seeder Transaction
		tx, err := db.Begin()
		if err != nil {
			log.Fatalf("[ERROR] %s", err)
		}

		defer tx.Rollback()

		qtx := queries.WithTx(tx)

		// Create "Chris"
		if err = qtx.SeedUser(ctx, sqlc.SeedUserParams{
			ID:       1,
			Uuid:     "3d7fdbb6-8b8d-490d-93fc-59c779fbc5c8",
			Name:     "Chris",
			Github:   sql.NullString{String: "https://github.com/chrislentz", Valid: true},
			Twitter:  sql.NullString{String: "https://twitter.com/ATLChris", Valid: true},
			Mastodon: sql.NullString{String: "https://mas.to/@ATLChris", Valid: true},
		}); err != nil {
			log.Fatalf("[ERROR] %s", err)
		}

		// Create "Golang"
		if err = qtx.SeedUser(ctx, sqlc.SeedUserParams{
			ID:      2,
			Uuid:    "9fd5f5af-ab29-46be-92be-65ade6b4d91c",
			Name:    "Golang",
			Github:  sql.NullString{String: "https://github.com/golang", Valid: true},
			Twitter: sql.NullString{String: "https://twitter.com/golang", Valid: true},
		}); err != nil {
			log.Fatalf("[ERROR] %s", err)
		}

		if err = tx.Commit(); err != nil {
			log.Fatalf("[ERROR] %s", err)
		}
	}

	return true
}
