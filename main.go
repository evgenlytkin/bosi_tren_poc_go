package main

import (
	"admissible_offer"
	"admissible_service"
	"base_travel_solution"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
	"travel_solution"

	"github.com/joho/godotenv"
)

// main function is the entry point of the application.
func main() {
	// Function calls commented out for potential future use.
	// one_call_with_one_big_response()
	// one_call_with_one_stubbed_big_response()
	// multiple_calls_with_response_in_parallel()
	// multiple_calls_with_multiple_small_stubbed_responses()
}

// one_call_with_one_big_response simulates a scenario where a single request is sent and a large response is expected.
func one_call_with_one_big_response() {
	var client = createHttpClient() // Create an HTTP client.

	send_first_three_request(client) // Send the first three requests.

	stubbed_payload := base_travel_solution.GetPayload() // Get the payload for the request.
	// Send the request and measure elapsed times for request and parsing.
	_, err, elapsed_request_time, elapsed_parsing_time := base_travel_solution.SendRequest(stubbed_payload, []byte(""), client)

	// Log the maximum elapsed times.
	log.Printf("max_elapsed_request_time: %f", elapsed_request_time)
	log.Printf("max_elapsed_parsing_time: %f", elapsed_parsing_time)

	if err != nil {
		fmt.Println(err) // Print the error if any.
		return
	}
}

// one_call_with_one_stubbed_big_response simulates a scenario with a stubbed response to avoid actual network calls.
func one_call_with_one_stubbed_big_response() {
	var client = createHttpClient()  // Create an HTTP client.
	send_first_three_request(client) // Send the first three requests.

	stubbed_payload := base_travel_solution.GetPayload()   // Get the stubbed payload.
	stubbed_response := base_travel_solution.GetResponse() // Get the stubbed response.
	// Send the request with stubbed data and measure elapsed times.
	_, err, elapsed_request_time, elapsed_parsing_time := base_travel_solution.SendRequest([]byte(stubbed_payload), []byte(stubbed_response), client)

	// Log the maximum elapsed times.
	log.Printf("max_elapsed_request_time: %f", elapsed_request_time)
	log.Printf("max_elapsed_parsing_time: %f", elapsed_parsing_time)

	if err != nil {
		fmt.Println(err) // Print the error if any.
		return
	}
}

// multiple_calls_with_response_in_parallel simulates sending multiple requests in parallel and processing their responses.
func multiple_calls_with_response_in_parallel() {
	var client = createHttpClient() // Create an HTTP client.

	send_first_three_request(client) // Send the first three requests.

	stubbed_payloads := base_travel_solution.GetMultiplePayloads() // Get multiple payloads for the requests.

	var wg sync.WaitGroup                 // Use a WaitGroup to wait for all goroutines to finish.
	semaphore := make(chan struct{}, 100) // Use a semaphore to limit concurrency.

	max_elapsed_request_time := float64(0) // Track the maximum elapsed request time.
	max_elapsed_parsing_time := float64(0) // Track the maximum elapsed parsing time.

	for _, stubbed_payload := range stubbed_payloads {
		wg.Add(1)               // Increment the WaitGroup counter.
		semaphore <- struct{}{} // Acquire a semaphore slot.
		go func(payload string, client *http.Client) {
			defer wg.Done()                // Decrement the WaitGroup counter when the goroutine completes.
			defer func() { <-semaphore }() // Release the semaphore slot.
			// Send the request and measure elapsed times.
			_, err, elapsed_request_time, elapsed_parsing_time := base_travel_solution.SendRequest([]byte(payload), []byte(""), client)

			maxDuration(&max_elapsed_request_time, elapsed_request_time) // Update the maximum elapsed request time.
			maxDuration(&max_elapsed_parsing_time, elapsed_parsing_time) // Update the maximum elapsed parsing time.

			if err != nil {
				fmt.Println(err) // Print the error if any.
				return
			}
		}(stubbed_payload, client)
	}
	wg.Wait() // Wait for all goroutines to finish.

	// Log the maximum elapsed times.
	log.Printf("max_elapsed_request_time: %f", max_elapsed_request_time)
	log.Printf("max_elapsed_parsing_time: %f", max_elapsed_parsing_time)
}

// multiple_calls_with_multiple_small_stubbed_responses simulates sending multiple requests with stubbed responses in parallel.
func multiple_calls_with_multiple_small_stubbed_responses() {
	var client = createHttpClient() // Create an HTTP client.

	send_first_three_request(client) // Send the first three requests.

	stubbed_payloads := base_travel_solution.GetMultiplePayloads()   // Get multiple payloads.
	stubbed_responses := base_travel_solution.GetMultipleResponses() // Get multiple stubbed responses.
	stubbed_data := zip(stubbed_payloads, stubbed_responses)         // Zip payloads and responses together.

	var wg sync.WaitGroup // Use a WaitGroup to wait for all goroutines to finish.

	max_elapsed_request_time := float64(0) // Track the maximum elapsed request time.
	max_elapsed_parsing_time := float64(0) // Track the maximum elapsed parsing time.

	for _, pair := range stubbed_data {
		wg.Add(1) // Increment the WaitGroup counter.
		go func(payload, response string) {
			// Send the request with stubbed data and measure elapsed times.
			_, err, elapsed_request_time, elapsed_parsing_time := base_travel_solution.SendRequest([]byte(payload), []byte(response), client)
			maxDuration(&max_elapsed_request_time, elapsed_request_time) // Update the maximum elapsed request time.
			maxDuration(&max_elapsed_parsing_time, elapsed_parsing_time) // Update the maximum elapsed parsing time.
			defer wg.Done()                                              // Decrement the WaitGroup counter when the goroutine completes.

			if err != nil {
				fmt.Println(err) // Print the error if any.
				return
			}
		}(pair[0], pair[1])
	}

	wg.Wait() // Wait for all goroutines to finish.

	// Log the maximum elapsed times.
	log.Printf("max_elapsed_request_time: %f", max_elapsed_request_time)
	log.Printf("max_elapsed_parsing_time: %f", max_elapsed_parsing_time)
}

// maxDuration updates the maximum value between two float64 numbers.
func maxDuration(max_value *float64, value float64) {
	if value > *max_value {
		*max_value = value // Update the maximum value if the new value is greater.
	}
}

// send_first_three_request sends the first three predefined requests.
func send_first_three_request(client *http.Client) {
	// Send a request and handle errors.
	_, err := travel_solution.SendRequest(client)
	if err != nil {
		fmt.Println(err) // Print the error if any.
		return
	}

	// Send another request and handle errors.
	_, err = admissible_service.SendRequest(client)
	if err != nil {
		fmt.Println(err) // Print the error if any.
		return
	}

	// Send the third request and handle errors.
	_, err = admissible_offer.SendRequest(client)
	if err != nil {
		fmt.Println(err) // Print the error if any.
		return
	}
}

// createHttpClient creates and configures an HTTP client with a proxy and timeouts.
func createHttpClient() *http.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	proxy_url := os.Getenv("PROXY")     // Define the proxy URL.
	proxyURL, _ := url.Parse(proxy_url) // Parse the proxy URL.
	return &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:    600,                     // Set the maximum number of idle connections.
			IdleConnTimeout: 600 * time.Second,       // Set the idle connection timeout.
			Proxy:           http.ProxyURL(proxyURL), // Set the proxy.
		},
		Timeout: 600 * time.Second, // Set the request timeout.
	}
}

// zip combines two slices of strings into a slice of string slices, pairing elements by their indexes.
func zip(a, b []string) [][]string {
	zipped := make([][]string, 0, len(a)) // Initialize the zipped slice with capacity equal to the length of slice a.
	for i := range a {
		if i < len(b) { // Ensure the index is within the bounds of slice b.
			zipped = append(zipped, []string{a[i], b[i]}) // Append the paired elements to the zipped slice.
		}
	}
	return zipped // Return the zipped slice.
}
