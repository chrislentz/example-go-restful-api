package main

import (
	"log"

	"github.com/chrislentz/example-go-restful-api/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Enable line numbers in logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Load .env variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("[Error] Failed to load .env file")
	}

	app := echo.New()

	app.Use(middleware.Logger())

	v1 := app.Group("/v1")

	v1.GET("/users", routes.GetUsers)
	v1.GET("/users/:user_id", routes.GetUser)

	app.Logger.Fatal(app.Start(":8080"))
}
