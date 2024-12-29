package main

import (
	"fmt"
	"log"
	"os"

	"fiber-url-shortener/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

// setupRoutes configures the API endpoints for the application.
// It defines two main routes:
// - GET "/:url": Resolves a shortened URL to the original URL and redirects the user.
// - POST "/api/v1": Accepts a URL from the client and returns a shortened version.
func setupRoutes(app *fiber.App) {
	// Route to resolve shortened URLs to their original destinations.
	app.Get("/:url", routes.ResolveURL)

	// Route to create a shortened URL from the provided original URL.
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	// Load environment variables from the .env file.
	err := godotenv.Load()
	if err != nil {
		// Log an error if the .env file could not be loaded, but continue execution.
		fmt.Println(err)
	}

	// Create a new Fiber app instance.
	app := fiber.New()

	// Use the logger middleware to log all requests for debugging and monitoring.
	app.Use(logger.New())

	// Set up the application routes.
	setupRoutes(app)

	// Start the Fiber server and listen on the port specified in the environment variable APP_PORT.
	// If the server fails to start, log the error and exit the program.
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
