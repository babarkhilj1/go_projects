package routes

import (
	"fiber-url-shortener/database"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

// ResolveURL handles the resolution of a shortened URL to its original URL.
// It queries the Redis database for the short identifier, redirects to the original URL if found,
// and increments a redirection counter for tracking usage.
func ResolveURL(c *fiber.Ctx) error {
	// Extract the short identifier from the URL parameter.
	url := c.Params("url")

	// Query the Redis database (DB 0) for the original URL associated with the short identifier.
	r := database.CreateClient(0) // Redis client for URL storage.
	defer r.Close()

	// Get the original URL from the database.
	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		// If the short identifier is not found in the database, return a 404 Not Found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "short not found on database",
		})
	} else if err != nil {
		// If there's an error connecting to the database, return a 500 Internal Server Error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot connect to DB",
		})
	}

	// Increment the redirection counter in Redis (DB 1) for analytics or tracking purposes.
	rInr := database.CreateClient(1) // Redis client for the counter.
	defer rInr.Close()
	_ = rInr.Incr(database.Ctx, "counter") // Increment the "counter" key.

	// Redirect the user to the original URL with a 301 Moved Permanently status.
	return c.Redirect(value, 301)
}
