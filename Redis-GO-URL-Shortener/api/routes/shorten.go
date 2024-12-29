package routes

import (
	"os"
	"strconv"
	"time"

	"fiber-url-shortener/database"
	"fiber-url-shortener/helpers"

	"github.com/asaskevich/govalidator"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// request represents the structure of the incoming JSON payload for shortening a URL.
type request struct {
	URL         string        `json:"url"`    // The original URL to be shortened.
	CustomShort string        `json:"short"`  // Optional custom short identifier for the URL.
	Expiry      time.Duration `json:"expiry"` // Expiry time for the shortened URL in hours.
}

// response represents the structure of the JSON payload returned to the client.
type response struct {
	URL             string        `json:"url"`              // The original URL.
	CustomShort     string        `json:"short"`            // The generated or custom short URL.
	Expiry          time.Duration `json:"expiry"`           // Expiry time of the shortened URL in hours.
	XRateRemaining  int           `json:"rate_limit"`       // Remaining requests in the current rate limit window.
	XRateLimitReset time.Duration `json:"rate_limit_reset"` // Time (in minutes) until the rate limit resets.
}

// ShortenURL handles the creation of shortened URLs.
// It validates the input, applies rate limiting, generates or validates custom short identifiers,
// and stores the mapping in a Redis database.
func ShortenURL(c *fiber.Ctx) error {
	// Parse the incoming JSON request body into the `request` struct.
	body := new(request)
	if err := c.BodyParser(&body); err != nil {
		// Return a 400 Bad Request error if the body cannot be parsed.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	// Rate limiting: Check if the user's IP has remaining quota.
	r2 := database.CreateClient(1) // Use Redis database 1 for rate limiting.
	defer r2.Close()
	val, err := r2.Get(database.Ctx, c.IP()).Result()
	if err == redis.Nil {
		// If the IP is not found in the database, initialize it with the API quota and expiry.
		_ = r2.Set(database.Ctx, c.IP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err()
	} else {
		// If the IP is found, check the remaining quota.
		val, _ = r2.Get(database.Ctx, c.IP()).Result()
		valInt, _ := strconv.Atoi(val)
		if valInt <= 0 {
			// If the quota is exhausted, return a rate limit exceeded error.
			limit, _ := r2.TTL(database.Ctx, c.IP()).Result()
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"error":            "Rate limit exceeded",
				"rate_limit_reset": limit / time.Nanosecond / time.Minute,
			})
		}
	}

	// Validate the provided URL.
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid URL",
		})
	}

	// Prevent shortening of the domain itself to avoid infinite redirect loops.
	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"error": "haha... nice try",
		})
	}

	// Enforce HTTPS for the provided URL.
	body.URL = helpers.EnforceHTTP(body.URL)

	// Generate a short identifier: Use a custom short ID if provided, else generate a random one.
	var id string
	if body.CustomShort == "" {
		id = uuid.New().String()[:6] // Generate a 6-character random ID.
	} else {
		id = body.CustomShort
	}

	// Check for collisions in the database for the short ID.
	r := database.CreateClient(0) // Use Redis database 0 for URL storage.
	defer r.Close()
	val, _ = r.Get(database.Ctx, id).Result()
	if val != "" {
		// If the short ID is already in use, return a conflict error.
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "URL short already in use",
		})
	}

	// Set the expiry for the shortened URL, defaulting to 24 hours if not provided.
	if body.Expiry == 0 {
		body.Expiry = 24
	}

	// Store the shortened URL in the database with its expiry time.
	err = r.Set(database.Ctx, id, body.URL, body.Expiry*3600*time.Second).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to connect to server",
		})
	}

	// Construct the response object with URL details and rate limit information.
	resp := response{
		URL:             body.URL,
		CustomShort:     "",
		Expiry:          body.Expiry,
		XRateRemaining:  10, // Default remaining quota.
		XRateLimitReset: 30, // Default reset time in minutes.
	}
	r2.Decr(database.Ctx, c.IP()) // Decrease the rate limit quota.
	val, _ = r2.Get(database.Ctx, c.IP()).Result()
	resp.XRateRemaining, _ = strconv.Atoi(val)
	ttl, _ := r2.TTL(database.Ctx, c.IP()).Result()
	resp.XRateLimitReset = ttl / time.Nanosecond / time.Minute

	// Include the shortened URL in the response.
	resp.CustomShort = os.Getenv("DOMAIN") + "/" + id

	// Return the response as JSON with a 200 OK status.
	return c.Status(fiber.StatusOK).JSON(resp)
}
