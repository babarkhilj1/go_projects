package api

import (
	"context"  // Provides functionality to handle deadlines, cancellations, and other context-aware tasks
	"fmt"      // Used for formatted I/O
	"net/http" // Provides HTTP client and server implementations
	"time"     // Used to manage timeouts and intervals

	"go-twilio-verify/data" // Importing the data package for models and data-related utilities

	"github.com/gin-gonic/gin" // Importing the Gin web framework
)

// appTimeout defines the maximum duration allowed for an operation before timing out
const appTimeout = time.Second * 10

// sendSMS handles the API endpoint for sending OTPs via SMS
func (app *Config) sendSMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a context with a timeout to ensure operations complete within a defined period
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		defer cancel() // Ensure context resources are released

		// Variable to hold the incoming request payload
		var payload data.OTPData

		// Validate the request body and bind it to the payload
		app.validateBody(c, &payload)

		// Create a new OTPData object using the validated payload
		newData := data.OTPData{
			PhoneNumber: payload.PhoneNumber,
		}

		// Call the Twilio service to send the OTP
		_, err := app.twilioSendOTP(newData.PhoneNumber)
		if err != nil {
			// Respond with an error if OTP sending fails
			app.errorJSON(c, err)
			return
		}

		// Respond with a success message if OTP is sent successfully
		app.writeJSON(c, http.StatusAccepted, "OTP sent successfully")
	}
}

// verifySMS handles the API endpoint for verifying OTPs sent via SMS
func (app *Config) verifySMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a context with a timeout to ensure operations complete within a defined period
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		defer cancel() // Ensure context resources are released

		// Variable to hold the incoming request payload
		var payload data.VerifyData

		// Validate the request body and bind it to the payload
		app.validateBody(c, &payload)

		// Create a new VerifyData object using the validated payload
		newData := data.VerifyData{
			User: payload.User,
			Code: payload.Code,
		}

		// Call the Twilio service to verify the OTP
		err := app.twilioVerifyOTP(newData.User.PhoneNumber, newData.Code)
		fmt.Println("err: ", err) // Log the error for debugging
		if err != nil {
			// Respond with an error if OTP verification fails
			app.errorJSON(c, err)
			return
		}

		// Respond with a success message if OTP is verified successfully
		app.writeJSON(c, http.StatusAccepted, "OTP verified successfully")
	}
}
