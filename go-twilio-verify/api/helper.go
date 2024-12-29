package api

import (
	"net/http" // Provides HTTP status codes and functions

	"github.com/gin-gonic/gin"               // Gin framework for HTTP handling
	"github.com/go-playground/validator/v10" // Validation library for struct-based validation
)

// jsonResponse defines the structure of JSON responses sent to the client
type jsonResponse struct {
	Status  int    `json:"status"`  // HTTP status code
	Message string `json:"message"` // Response message (e.g., success or error message)
	Data    any    `json:"data"`    // Response data, flexible to hold any type
}

// Create a new instance of the validator library
var validate = validator.New()

// validateBody validates the incoming request body
func (app *Config) validateBody(c *gin.Context, data any) error {
	// Bind JSON data from the request body to the provided struct
	if err := c.BindJSON(&data); err != nil {
		return err // Return an error if binding fails
	}
	// Use the validator library to validate the struct fields
	if err := validate.Struct(&data); err != nil {
		return err // Return an error if validation fails
	}

	return nil // Return nil if validation is successful
}

// writeJSON sends a success response to the client with the given status and data
func (app *Config) writeJSON(c *gin.Context, status int, data any) {
	// Format the response using the jsonResponse structure
	c.JSON(status, jsonResponse{Status: status, Message: "success", Data: data})
}

// errorJSON sends an error response to the client with the given error message and optional status code
func (app *Config) errorJSON(c *gin.Context, err error, status ...int) {
	// Default to HTTP 400 Bad Request if no status code is provided
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	// Format the response using the jsonResponse structure with the error message
	c.JSON(statusCode, jsonResponse{Status: statusCode, Message: err.Error()})
}
