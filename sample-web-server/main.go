package main

import (
	"fmt"
	"net/http"
	"os"
)

var version = "v1.0.0"

func main() {
	// Print application version
	fmt.Println("Version: ", version)

	// Register handler functions for different endpoints
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/healthz", handleHealthz)

	// Start the server
	fmt.Println("Starting the server at port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
		os.Exit(1)
	}
}

// helloHandler is a simple HTTP handler function that writes a response with the hostname of the machine.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Fprint(w, "Error getting hostname")
		return
	}
	fmt.Fprintf(w, "Hello from %s. Version: %s\n", hostname, version)
}

// handleHealthz is a simple HTTP handler function that check the status of application and response.
func handleHealthz(w http.ResponseWriter, r *http.Request) {
	// it will check the application configuration here
	// if everything is fine then it will return OK
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}
