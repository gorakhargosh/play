// +build appengine

package main

import (
	"log"
	"net/http"
	"sync"
)

var (
	initServerOnce sync.Once
	server         http.Handler
)

// initServer initializes a server instance.
func initServer() {
	s, err := NewServer(Config{
		Debug: true,

		// TODO(yesudeep): make runtime.GOROOT()-relative.
		TemplatesPath: "templates",
	})
	if err != nil {
		log.Fatal(err)
	}
	server = &ErrorHandler{s}
}

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		initServerOnce.Do(initServer)
		server.ServeHTTP(w, r)
	})
}
