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
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

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
