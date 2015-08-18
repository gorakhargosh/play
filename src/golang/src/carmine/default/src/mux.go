package main

import "net/http"

func getMux(c Config) http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", &templateHandler{
		debug:         c.Debug,
		filenames:     []string{"index.html"},
		templatesPath: c.TemplatesPath,
	})
	return mux
}
