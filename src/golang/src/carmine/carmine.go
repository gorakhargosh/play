// HTTP server implementation for the Blood emergency application.
//
// +build !appengine

package main

import (
	"fmt"
	"log"
	"net/http"
)

// Holds the configuration for the application.
var config Config

func main() {
	hostAddr := fmt.Sprintf("%s:%d", config.Host, config.Port)

	mux := http.NewServeMux()
	mux.Handle("/", &templateHandler{
		debug:        config.Debug,
		filenames:    []string{"index.html"},
		templatesDir: config.TemplatesDir,
	})
	log.Printf("starting HTTP server on %s...\n", hostAddr)
	err := http.ListenAndServe(hostAddr, mux)
	if err != nil {
		log.Fatalf("cannot start HTTP server: error - %s", err)
	}
}
