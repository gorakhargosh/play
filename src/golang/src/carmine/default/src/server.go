package main

import (
	"html/template"
	"net/http"
)

// Server implements a carmine application server.
type Server struct {
	config   Config
	template struct {
		index *template.Template
	}
}

// NewServer constructs a new server using the specified configuration.
func NewServer(config Config) (*Server, error) {
	s := &Server{config: config}
	s.template.index = parse(config.TemplatesPath, "base.html", "index.html")
	return s, nil
}

// NotFound writes as error resonse: 404.
func (s *Server) NotFound(w http.ResponseWriter, r *http.Request) {
	// TODO(yesudeep): Replace this with a custom template rendering later.
	http.NotFound(w, r)
}

// ServeHTTP implements the http.Handler interface.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var t *template.Template

	switch r.URL.Path {
	case "/":
		t = s.template.index
		w.Header().Set("Content-Type", "text/html")
	default:
		s.NotFound(w, r)
	}

	err := t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
