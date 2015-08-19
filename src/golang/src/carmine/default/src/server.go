package main

import (
	"html/template"
	"net/http"
)

// Server implements the application server.
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

// ServeHTTP implements the HTTPErrorHandler interface.
func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	var t *template.Template
	contentType := "text/html"
	switch r.URL.Path {
	case "/":
		t = s.template.index
	default:
		return NewHTTPError(http.StatusNotFound, "Not found")
	}
	w.Header().Set("Content-Type", contentType)
	return t.Execute(w, nil)
}
