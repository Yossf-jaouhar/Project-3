package main

import (
	"fmt"
	"log"
	"net/http"

	"ascii/funcs"
)

func main() {
	// Set up the file server to serve static assets from the "assets" directory.
	// The static files (CSS,images) are accessible via /assets/.
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Set up the HTTP handlers for the root and ASCII art routes.
	http.HandleFunc("/", funcs.HomeHandler)
	http.HandleFunc("/ascii-art", funcs.AsciiArtHandler)
	http.HandleFunc("/Download", funcs.Download)
	// Define the server address.
	port := ":8080"
	fmt.Printf("Server is running at http://localhost%s/\n", port)

	// Start the server and log any errors.
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
