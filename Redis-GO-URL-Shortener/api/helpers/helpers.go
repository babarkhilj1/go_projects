package helpers

import (
	"os"
	"strings"
)

// EnforceHTTP ensures that a URL starts with "http://" or "https://".
// If the URL does not have a scheme, it prepends "http://" by default.
func EnforceHTTP(url string) string {
	// Check if the URL does not start with "http" and prepend "http://".
	if url[:4] != "http" {
		return "http://" + url
	}
	// If the URL already has a valid scheme, return it as is.
	return url
}

// RemoveDomainError validates that the given URL does not match the application's domain.
// This is used to prevent users from creating shortened URLs that point to the application's domain itself,
// which could cause infinite redirect loops.
func RemoveDomainError(url string) bool {
	// Check if the provided URL matches the domain specified in the environment variable.
	if url == os.Getenv("DOMAIN") {
		return false
	}

	// Remove common prefixes like "http://", "https://", and "www." from the URL.
	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)

	// Extract the base domain by splitting the URL at the first "/" after the domain.
	newURL = strings.Split(newURL, "/")[0]

	// If the processed URL matches the application's domain, return false to indicate an error.
	if newURL == os.Getenv("DOMAIN") {
		return false
	}

	// Return true if the URL does not conflict with the application's domain.
	return true
}
