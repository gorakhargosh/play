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

type Config struct {
	Host string
	Port int
}

var config Config

func init() {
	config = Config{}
	flag.StringVar(&config.Host, "host", "", "the host on which to listen")
	flag.IntVar(&config.Port, "port", 8080, "the port on which to listen")
	flag.Parse()
}

type baseHandler struct {
}

func (h baseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func main() {
	hostAddr := fmt.Sprintf("%s:%d", config.Host, config.Port)

	mux := http.NewServeMux()
	mux.Handle("/", &baseHandler{})
	log.Printf("Starting HTTP server on %s...\n", hostAddr)
	err := http.ListenAndServe(hostAddr, mux)
	if err != nil {
		log.Fatalf("Could not start HTTP server: error - %s", err)
	}
}
