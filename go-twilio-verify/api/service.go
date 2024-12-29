package api

import (
	"errors" // Used to handle error messages

	"github.com/twilio/twilio-go"                          // Twilio SDK for interacting with Twilio services
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2" // Specific Verify API package from Twilio
)

// Initialize Twilio client with account credentials
var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: envACCOUNTSID(), // Twilio Account SID from environment variables
	Password: envAUTHTOKEN(),  // Twilio Auth Token from environment variables
})

// twilioSendOTP sends an OTP to the provided phone number via SMS
func (app *Config) twilioSendOTP(phoneNumber string) (string, error) {
	// Set up parameters for the verification request
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(phoneNumber) // Set the recipient phone number
	params.SetChannel("sms")  // Specify SMS as the delivery channel

	// Make a request to Twilio's Verify API to create a verification
	resp, err := client.VerifyV2.CreateVerification(envSERVICESID(), params)
	if err != nil {
		return "", err // Return an error if the API call fails
	}

	return *resp.Sid, nil // Return the verification SID on success
}

// twilioVerifyOTP verifies the OTP sent to the provided phone number
func (app *Config) twilioVerifyOTP(phoneNumber string, code string) error {
	// Set up parameters for the verification check
	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo(phoneNumber) // Set the recipient phone number
	params.SetCode(code)      // Set the OTP code to verify

	// Make a request to Twilio's Verify API to check the OTP
	resp, err := client.VerifyV2.CreateVerificationCheck(envSERVICESID(), params)
	if err != nil {
		return err // Return an error if the API call fails
	}

	// Twilio Verify API returns a status. If it's not "approved", the OTP is invalid.
	if *resp.Status != "approved" {
		return errors.New("not a valid code") // Return an error for invalid OTP
	}

	return nil // Return nil if the OTP is verified successfully
}
