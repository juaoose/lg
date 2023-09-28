package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type Application struct {
	logger *slog.Logger
}

func main() {
	// Retrieve command line attributes
	addr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()

	app := &Application{
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}

	// Start server
	app.logger.Info("starting server", "address", *addr)
	err := http.ListenAndServe(*addr, app.routes())
	app.logger.Error(err.Error())
	os.Exit(1)
}
