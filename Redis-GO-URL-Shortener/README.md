# Redis Go URL Shortener

A lightweight URL shortener built with Go, Fiber, and Redis. This project enables users to shorten URLs, resolve shortened URLs back to their original form, and implements rate limiting for API requests.

---

## Features

- **URL Shortening**: Convert long URLs into short, easily shareable links.
- **URL Redirection**: Automatically redirect users from the short link to the original URL.
- **Rate Limiting**: Limit API usage to prevent abuse (default: 10 requests per 30 minutes).
- **Custom Short URLs**: Users can provide their own custom short codes.
- **Redis Database**: Uses Redis for fast and efficient storage of URLs and request metadata.
- **Dockerized**: Fully containerized using Docker for easy deployment.

---

## Project Structure

```plaintext
|   docker-compose.yml
|
+---api
|   |   .env
|   |   Dockerfile
|   |   go.mod
|   |   go.sum
|   |   main.go
|   |
|   +---database
|   |       database.go
|   |
|   +---helpers
|   |       helpers.go
|   |
|   \---routes
|           resolve.go
|           shorten.go
|
+---data
\---db
        Dockerfile
```

---

## Prerequisites

- Docker and Docker Compose installed on your system.
- A `.env` file in the `api` directory with the following variables:
  ```dotenv
  DOMAIN=localhost:3000
  APP_PORT=:3000
  DB_ADDR=redis:6379
  DB_PASS=  # Leave empty if no Redis password is set
  API_QUOTA=10
  ```

---

## Installation

1. **Run the application using Docker Compose**:
   ```bash
   docker-compose up --build
   ```

2. **Access the application**:
   - Shorten URLs: `POST http://localhost:3000/api/v1`
   - Resolve URLs: Open `http://localhost:3000/{short_code}` in your browser.

---

## API Endpoints

### 1. Shorten URL
**Endpoint**: `POST /api/v1`  
**Request Body**:
```json
{
  "url": "https://example.com",
  "short": "customShortCode", // Optional
  "expiry": 24 // Expiry in hours (default: 24)
}
```

**Response**:
```json
{
  "url": "https://example.com",
  "short": "http://localhost:3000/customShortCode",
  "expiry": 24,
  "rate_limit": 9,
  "rate_limit_reset": 30
}
```

### 2. Resolve URL
**Endpoint**: `GET /{short_code}`  
- Redirects to the original URL if the short code exists.

**Error Response**:
```json
{
  "error": "short not found on database"
}
```

---

## Files Explained

- **`api/main.go`**: Entry point for the application, initializes routes and middleware.
- **`api/routes/shorten.go`**: Handles the logic for shortening URLs and applying rate limits.
- **`api/routes/resolve.go`**: Handles resolving short URLs back to their original form.
- **`api/helpers/helpers.go`**: Contains utility functions for URL validation and manipulation.
- **`api/database/database.go`**: Provides a Redis client for database interactions.
- **`api/Dockerfile`**: Docker configuration for the API service.
- **`db/Dockerfile`**: Docker configuration for the Redis service.
- **`docker-compose.yml`**: Orchestrates the API and Redis services using Docker Compose.

---

## Testing the Application

1. Use a tool like [Postman](https://www.postman.com/) or `curl` for API testing.
2. Example commands:
   - Shorten a URL:
     ```bash
     curl -X POST http://localhost:3000/api/v1 \
     -H "Content-Type: application/json" \
     -d '{"url": "https://example.com", "expiry": 24}'
     ```
   - Resolve a URL:
     ```bash
     curl -X GET http://localhost:3000/{short_code}
     ```

---

## Future Enhancements

- Add collision detection for auto-generated short codes.
- Support for custom domain configuration.
- Admin dashboard for managing shortened URLs.
- Enhanced analytics and tracking for short URLs.

---