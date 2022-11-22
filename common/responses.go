package common

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// Response codes
type ResponseCode struct {
	Message    string
	HttpStatus int
}

var response_codes = map[string]ResponseCode{
	// Miscellaneous Codes
	"m-0001": {
		Message:    "An unknown error has occured.",
		HttpStatus: http.StatusInternalServerError,
	},

	//  User Codes
	"u-0001": {
		Message:    "The given user id is invalid.",
		HttpStatus: http.StatusUnprocessableEntity,
	},
}

// Error responses
type ErrorObject struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ResponseError struct {
	Error ErrorObject `json:"error"`
}

func RespondUnknown(c echo.Context, err error) error {
	log.Printf("[ERROR] %s", err)

	return RespondCode(c, "m-0001")
}

func RespondCode(c echo.Context, code string) error {
	response := ResponseError{ErrorObject{Code: code, Message: response_codes[code].Message}}
	return c.JSON(response_codes[code].HttpStatus, response)
}

// Successful responses
type ResponseItem struct {
	Data interface{} `json:"data"`
}

func (r *ResponseItem) Respond(c echo.Context) error {
	return c.JSON(http.StatusOK, r.Data)
}

type ResponseCollection struct {
	PreviousPage string      `json:"previous_page"`
	NextPage     string      `json:"next_page"`
	Data         interface{} `json:"data"`
}

func (r *ResponseCollection) Respond(c echo.Context) error {
	return c.JSON(http.StatusOK, r)
}

type ResponseOperation struct {
	Success bool        `json:"success"`
	Meta    interface{} `json:"meta,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (r *ResponseOperation) RespondCreated(c echo.Context) {
	r.Success = true
	c.JSON(http.StatusCreated, r)
}

func (r *ResponseOperation) RespondOK(c echo.Context) {
	r.Success = true
	c.JSON(http.StatusOK, r)
}
