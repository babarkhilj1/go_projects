package api

import (
	"log" // Used for logging errors and messages
	"os"  // Provides functions to interact with the operating system, like environment variables

	"github.com/joho/godotenv" // Library for loading environment variables from a .env file
)

// envACCOUNTSID retrieves the Twilio Account SID from the .env file
func envACCOUNTSID() string {
	// Load environment variables from the .env file
	err := godotenv.Load(".env")
	if err != nil {
		// Log the error and terminate the application if .env cannot be loaded
		log.Fatalln(err)
		log.Fatal("Error loading .env file")
	}
	// Return the TWILIO_ACCOUNT_SID value from the environment variables
	return os.Getenv("TWILIO_ACCOUNT_SID")
}

// envAUTHTOKEN retrieves the Twilio Auth Token from the .env file
func envAUTHTOKEN() string {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		// Log the error and terminate the application if .env cannot be loaded
		log.Fatal("Error loading .env file")
	}
	// Return the TWILIO_AUTHTOKEN value from the environment variables
	return os.Getenv("TWILIO_AUTHTOKEN")
}

// envSERVICESID retrieves the Twilio Service SID from the .env file
func envSERVICESID() string {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		// Log the error and terminate the application if .env cannot be loaded
		log.Fatal("Error loading .env file")
	}
	// Return the TWILIO_SERVICES_ID value from the environment variables
	return os.Getenv("TWILIO_SERVICES_ID")
}
