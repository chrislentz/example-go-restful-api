package routes

import (
	"context"

	"github.com/chrislentz/example-go-restful-api/common"
	"github.com/chrislentz/example-go-restful-api/database"
	"github.com/chrislentz/example-go-restful-api/sqlc"
	"github.com/chrislentz/example-go-restful-api/transformers"
	"github.com/labstack/echo"
)

// Name: Get user
// Request Method: GET
// Path: /user/:user_id
// Documentation: https://github.com/chrislentz/example-go-restful-api-bash-history
func GetUser(c echo.Context) error {
	var response common.ResponseItem

	// Path params
	userId := c.Param("user_id")

	db := database.Connect()

	queries := sqlc.New(db)
	ctx := context.Background()

	// Lookup user instance
	user, err := queries.GetUserByID(ctx, userId)
	if err != nil {
		return common.RespondCode(c, "u-0001")
	}

	// Transform and return response
	transformer := transformers.UserTransformer{User: user}
	response.Data = transformer.Transform()
	return response.Respond(c)
}
