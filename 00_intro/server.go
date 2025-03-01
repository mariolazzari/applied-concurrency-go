package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Configure path and handler function
	http.HandleFunc("/hello", helloHandler)
	// Listen on port 8080 and block main
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Write response
	fmt.Fprintf(w, "Hello, Mario!")
}
