// HTTP server implementation for the Blood emergency application.
//
// +build !appengine

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

// The HTTP server mux.
var mux http.ServeMux

// Holds application configuration.
var config Config

func init() {
	config = Config{}
	flag.StringVar(&config.Host, "host", "", "the host on which to listen")
	flag.IntVar(&config.Port, "port", 8080, "the port on which to listen")
	flag.BoolVar(&config.Debug, "debug", false, "turns on debugging")
	flag.StringVar(
		&config.TemplatesPath,
		"templatesDir",
		"templates", "the directory that contains the templates")
	flag.Parse()
	mux = getMux(config)
}

func main() {
	hostAddr := fmt.Sprintf("%s:%d", config.Host, config.Port)

	log.Printf("starting HTTP server on %s...\n", hostAddr)
	err := http.ListenAndServe(hostAddr, mux)
	if err != nil {
		log.Fatalf("cannot start HTTP server: error - %s", err)
	}
}
