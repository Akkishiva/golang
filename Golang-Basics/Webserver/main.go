package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)    // Register handler function for root path ("/")
	http.HandleFunc("/help", helphandler) // Register handler function for "/help"

	log.Printf("Starting server on http://localhost:8000\n")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the path component of the request URL r
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// helphandler provides help message for requests
func helphandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "helping the request")
}
