package api

import "github.com/gin-gonic/gin"

// Config defines the configuration structure for the application, including the router
type Config struct {
	Router *gin.Engine // Gin engine for routing
}

// Routes sets up the API endpoints for the application
func (app *Config) Routes() {
	// Define a POST route for sending OTPs
	app.Router.POST("/otp", app.sendSMS())

	// Define a POST route for verifying OTPs
	app.Router.POST("/verifyOTP", app.verifySMS())
}
