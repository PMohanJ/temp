package main

import (
	"fmt"
	"net/http"
)

// handler function for the root path
func handler(w http.ResponseWriter, r *http.Request) {
	// Writing "OK" to the response body
	fmt.Fprintf(w, "OK")
}

func main() {
	// Define the server address (IP:Port)
	address := "0.0.0.0:8080" // or use "localhost:8080" for local testing

	// Register the handler for the root route "/"
	http.HandleFunc("/", handler)

	// Start the server
	fmt.Println("Server is listening on", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}