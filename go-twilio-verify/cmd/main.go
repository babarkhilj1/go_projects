package main

import (
	"go-twilio-verify/api" // Importing the API package containing the app configuration and routes

	"github.com/gin-gonic/gin" // Importing the Gin web framework
)

func main() {
	// Create a new Gin router instance
	router := gin.Default()

	// Initialize application configuration
	// The Config struct from the api package is used to set up the router and other configurations
	app := api.Config{Router: router}

	// Set up application routes
	app.Routes()

	// Start the server on port 8000
	router.Run(":8000")
}
