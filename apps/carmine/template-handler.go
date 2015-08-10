package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"sync"
)

func compileHTMLTemplate(dir, filename string) *template.Template {
	return template.Must(template.ParseFiles(filepath.Join(dir, filename)))
}

type templateHandler struct {
	// Enabling debug mode causes the template to be recompiled per request
	// (useful during development).
	debug bool
	once  sync.Once

	filename     string
	templatesDir string
	template     *template.Template
}

func (h *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.debug {
		h.template = compileHTMLTemplate(h.templatesDir, h.filename)
	} else {
		h.once.Do(func() {
			h.template = compileHTMLTemplate(h.templatesDir, h.filename)
		})
	}
	h.template.Execute(w, r)
}
