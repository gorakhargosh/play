package main

import (
	"html/template"
	"log"
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
	s := &Server{
		config: config,
	}

	s.template.index = parse(config.TemplatesPath, "base.html", "index.html")
	return s, nil
}

// ServeHTTPError handles HTTP errors by rendering custom pages.
func (s *Server) ServeHTTPError(w http.ResponseWriter, r *http.Request, status int, err string) {
	w.WriteHeader(status)
	switch status {
	case http.StatusNotFound:
		// render custom template.
	case http.StatusInternalServerError:
		// render custom template.
	}
}

// ServeHTTP implements the http.Handler interface.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var t *template.Template

	switch p := r.URL.Path; p {
	case "/":
		t = s.template.index
		w.Header().Set("Content-Type", "text/html")
	default:
		s.ServeHTTPError(w, r, http.StatusNotFound, "Not found")
	}

	err := t.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
