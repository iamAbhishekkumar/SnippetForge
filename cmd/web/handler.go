package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
} // Add a snippetView handler function.

func viewSnippet(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string and try to
	// convert it to an integer
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Add a createSnippet handler function.
func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// If it's not, use the w.WriteHeader() method to send a 405 status
		// code and the w.Write() method to write a "Method Not Allowed"
		// response body.

		/*
			w.WriteHeader(405)
			w.Write([]byte("Method Not Allowed"))
			This can be also written as :
		*/
		// w.Header()["Date"] = nil // for suppressing any header
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}
