package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

var version = "v4.0.0"

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	// Print application version
	fmt.Println("Version: ", version)

	// Register handler functions for different endpoints
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/healthz", handleHealthz)
	http.HandleFunc("/books", handleBooks)

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

// handleBooks is a simple HTTP handler function that does the following:
// 1. When a POST request is sent with the book information as JSON, it adds the book to the list and store as a JSON file.
// 2. When a GET request is sent, it read the JSON file returns the list of books.
func handleBooks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Read the JSON file and return the list of books
		books, err := readBooks()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}
		booksBytes, err := json.MarshalIndent(books, "", "  ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(booksBytes))
	case http.MethodPost:
		// read the json payload from request
		var book Book
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, err.Error())
			return
		}
		// Add the book to the list and store it in a JSON file
		books, err := readBooks()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}
		// add the book to the list
		books = append(books, book)
		// save the new list to the JSON file
		err = saveBooks(books)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "Book added successfully")
	}
}

func readBooks() ([]Book, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(filepath.Join(currentDir, "data", "books.json"))
	if err != nil {
		// return empty list if file does not exist
		if os.IsNotExist(err) {
			return []Book{}, nil
		}
		return nil, err
	}
	var books []Book
	err = json.Unmarshal(data, &books)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func saveBooks(books []Book) error {
	bookBytes, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		return err
	}
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	filename := filepath.Join(currentDir, "data", "books.json")

	// If the file does not exist, create it
	if _, err := os.Stat(filepath.Dir(filename)); os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(filename), os.ModePerm)
		if err != nil {
			return err
		}
	}

	err = os.WriteFile(filename, bookBytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
