package main

import "net/http"

func (app *Application) routes() *http.ServeMux {
	/*
	 Register handlers: note that these functions are not Handlers
	 (they do not satisfy the http.Handler interface), but with mux.HandleFunc
	 we skip the requirement
	*/
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	// Serve static content
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
