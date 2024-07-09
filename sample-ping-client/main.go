package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {

	// Read the ping URL and ping interval from environment variables
	pingURL := os.Getenv("PING_URL")
	pingInterval := os.Getenv("PING_INTERVAL")

	// Exit if the server URL is not provided
	if pingURL == "" {
		panic("PING_URL environment variable is not provided")
	}

	// If the ping interval is not provided, default to 1 second
	if pingInterval == "" {
		pingInterval = "1s"
	}

	// Parse the string formatted ping interval to a time.Duration
	sleepDuration, err := time.ParseDuration(pingInterval)
	if err != nil {
		panic(err)
	}

	// Enter into an infinite loop to ping the server
	for {
		pingServer(pingURL)

		// Sleep for the specified interval
		time.Sleep(sleepDuration)
	}
}

func pingServer(serverURL string) {
	// Ping the server
	resp, err := http.Get(serverURL)
	if err != nil {
		fmt.Printf("Error pinging the server: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Check if the server returned an error
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Server returned an error: %v\n", resp.Status)
		return
	}

	// Read the response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading the server response: %v\n", err)
		return
	}

	// Print the response
	fmt.Printf("Server response: %v\n", string(data))
}
