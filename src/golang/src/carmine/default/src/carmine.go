// +build appengine

package main

import (
	"log"
	"net/http"
	"sync"
)

var (
	serverInitOnce sync.Once
	server         http.Handler
)

// serverInit initializes a server instance.
func serverInit() {
	// TODO(yesudeep): Determine the goroot using the runtime and use that here to
	// determine the templates directory filepath.
	s, err := NewServer(Config{
		Debug:         true,
		TemplatesPath: "templates",
	})
	if err != nil {
		log.Fatal(err)
	}
	server = s
}

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serverInitOnce.Do(serverInit)
		server.ServeHTTP(w, r)
	})
}
