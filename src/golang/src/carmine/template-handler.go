package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"sync"
)

const (
	leftDelim  = "{{{"
	rightDelim = "}}}"
)

// compileHTMLTemplate compiles a given sequence of templates for the
// application.
func compileHTMLTemplate(dir string, filenames ...string) *template.Template {
	paths := []string{}
	for _, filename := range filenames {
		paths = append(paths, filepath.Join(dir, filename))
	}
	return template.Must(template.ParseFiles(paths...))
}

type templateHandler struct {
	// Enabling debug mode causes the template to be recompiled per request
	// (useful during development).
	debug bool
	once  sync.Once

	filenames    []string
	templatesDir string
	template     *template.Template
}

func (h *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.debug {
		h.template = compileHTMLTemplate(h.templatesDir, h.filenames...)
	} else {
		h.once.Do(func() {
			h.template = compileHTMLTemplate(h.templatesDir, h.filenames...)
		})
	}
	h.template.Execute(w, r)
}
