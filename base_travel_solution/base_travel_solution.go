// Package base_travel_solution provides functionalities to send HTTP requests
// and convert XML responses to JSON format for travel solution services.
package base_travel_solution

import (
	"bytes"    // Import for working with bytes.
	"log"      // Import for logging errors.
	"net/http" // Import for HTTP client and request functionalities.
	"os"       // Import for accessing environment variables.
	"time"     // Import for measuring execution time.

	"utils" // Import custom utility package for HTTP requests.

	xml2json "github.com/basgys/goxml2json" // Import for converting XML to JSON.
)

// SendRequest sends an HTTP request to the BASE_TRAVEL_SOLUTION_ENDPOINT,
// optionally using stubbed responses, and returns the JSON response, error if any,
// and the elapsed time for the request and parsing.
func SendRequest(stubbed_payload []byte, stubbed_response []byte, client *http.Client) (string, error, float64, float64) {
	// Retrieve the endpoint URL from environment variables.
	endpoint := os.Getenv("BASE_TRAVEL_SOLUTION_ENDPOINT")
	// Define the HTTP method for the request.
	method := "POST"
	// Initialize request headers with content type and IBM Cloud credentials.
	headers := map[string]string{
		"Content-Type":        "text/xml",
		"X-IBM-Client-Id":     os.Getenv("X_IBM_CLIENT_ID"),
		"X-IBM-Client-Secret": os.Getenv("X_IBM_CLIENT_SECRET"),
	}
	var respBody []byte // Variable to store the response body.

	// Record the start time of the request for measuring duration.
	start_request_time := time.Now()

	// Check if a stubbed response is provided. If not, send the actual HTTP request.
	if len(stubbed_response) == 0 {
		body := stubbed_payload // Use the provided stubbed payload as the request body.

		// Send the HTTP request and ignore errors for simplicity.
		respBody, _ = utils.SendHTTPRequest(client, method, endpoint, body, headers)
	} else {
		// Use the provided stubbed response as the response body.
		respBody = stubbed_response
	}

	// Calculate the elapsed time for the request.
	elapsed_request_time := time.Since(start_request_time)

	// Record the start time of the parsing process for measuring duration.
	start_parsing_time := time.Now()
	// Create a reader from the response body for conversion.
	reader := bytes.NewReader(respBody)
	// Convert the XML response body to JSON format.
	jsonBuffer, err := xml2json.Convert(reader)
	if err != nil {
		// Log an error if the conversion fails.
		log.Printf("Error converting response: %v", err)
	}
	// Calculate the elapsed time for the parsing process.
	elapsed_parsing_time := time.Since(start_parsing_time)

	// Return the JSON response as a string, nil for error placeholder,
	// and the elapsed times for request and parsing.
	return jsonBuffer.String(), nil, elapsed_request_time.Seconds(), elapsed_parsing_time.Seconds()
}
