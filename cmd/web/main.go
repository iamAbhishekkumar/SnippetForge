package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Create a file server which serves files out of the "./ui/static" directory.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	/*
		Why do we need to remove the prefix :

		Suppose you have a file named "style.css" located at "./ui/static/css/style.css" on your server
		filesystem. You want this file to be accessible via a web browser using the URL path "/static/css/style.css".

		However, when you set up the file server, you tell it to serve files from the directory "./ui/static/".
		This means that when a request comes in for "/static/css/style.css", the file server will look for the file
		"./ui/static/static/css/style.css", which doesn't exist. The server is expecting the URL path to directly map
		 to the file structure within the "./ui/static/" directory.
	*/

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Register the two new handler functions and corresponding URL patterns with
	// the servemux, in exactly the same way that we did before.
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", viewSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
