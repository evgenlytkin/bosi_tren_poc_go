// Package utils contains utility functions for common operations.
package utils

import (
	"bytes"     // Import for working with bytes.
	"fmt"       // Import for formatting strings.
	"io/ioutil" // Import for reading from and writing to files.
	"log"       // Import for logging errors and information.
	"net/http"  // Import for making HTTP requests.
)

// SendHTTPRequest sends an HTTP request and returns the response body as a byte slice.
// It requires an http.Client for making the request, the HTTP method and endpoint as strings,
// the request body as a byte slice, and a map of headers. It returns the response body as a byte slice
// and an error if any occurred during the request process.
func SendHTTPRequest(client *http.Client, method, endpoint string, body []byte, headers map[string]string) ([]byte, error) {
	// Create a new HTTP request with the provided method, endpoint, and body.
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(body))
	if err != nil {
		// Log and return the error if the request creation fails.
		log.Printf("Error creating new request: %v", err)
		return []byte(""), err
	}

	// Set the provided headers on the request.
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Send the request using the provided http.Client.
	resp, err := client.Do(req)
	if err != nil {
		// Log and return the error if sending the request fails.
		log.Printf("Error sending request: %v", err)
		return []byte(""), err
	}

	// Check if the response status code is 403 Forbidden.
	if resp.StatusCode == http.StatusForbidden {
		// Return an error indicating access was forbidden.
		return []byte(""), fmt.Errorf("access forbidden: received 403 error for request to %s, status: %s", resp.Request.URL, resp.Status)
	}

	// Ensure the response body is closed after reading.
	defer resp.Body.Close()

	// Read the response body.
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Log and return the error if reading the response body fails.
		log.Printf("Error reading response body: %v", err)
		return []byte(""), err
	}

	// Return the response body and nil for error if the process was successful.
	return respBody, nil
}
