package routes

import (
	"context"

	"github.com/chrislentz/example-go-restful-api/common"
	"github.com/chrislentz/example-go-restful-api/database"
	"github.com/chrislentz/example-go-restful-api/sqlc"
	"github.com/chrislentz/example-go-restful-api/transformers"
	"github.com/labstack/echo"
)

// Name: List users
// Request Method: GET
// Path: /users
// Documentation: https://github.com/chrislentz/example-go-restful-api-bash-history
func GetUsers(c echo.Context) error {
	var response common.ResponseCollection

	db := database.Connect()

	queries := sqlc.New(db)
	ctx := context.Background()

	// Lookup all users
	users, err := queries.GetUsers(ctx)
	if err != nil {
		return common.RespondUnknown(c, err)
	}

	// Transform and return response
	transformer := transformers.UsersTransformer{Users: users}
	response.Data = transformer.Transform()
	return response.Respond(c)
}
