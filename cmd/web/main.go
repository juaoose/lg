package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Register handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Serve static content
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Retrieve command line attributes
	addr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()

	// Start server
	log.Printf("Starting server on port %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
