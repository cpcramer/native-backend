package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Set up a simple HTTP server that listens on port 8080
	fmt.Printf("Starting server...\n")

	fmt.Printf("listening...")

	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	}))
}
