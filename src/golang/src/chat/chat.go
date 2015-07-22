package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

// Turns on debugging for the application; especially useful during
// development:
//
// 1. Templates are compiled per request instead of per process instance
//    allowing reloading the templates as you make changes.
var debug = flag.Bool("debug", false, "turns on debugging")

// The port to which the HTTP server will bind.
var listenPort = flag.Int("port", 80, "listens on port number")

// compileTemplate compiles a template given its filename.
func compileHTMLTemplate(filename string) *template.Template {
	return template.Must(template.ParseFiles(filepath.Join("templates", filename)))
}

// A templateHandler executes a template to generate an HTTP response.
type templateHandler struct {
	// Enabling debug mode causes the template to be recompiled
	// per request (useful during development).
	debug bool
	once  sync.Once

	filename string
	templ    *template.Template
}

// ServeHTTP implements the http.Handler interface.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if t.debug {
		t.templ = compileHTMLTemplate(t.filename)
	} else {
		t.once.Do(func() {
			t.templ = compileHTMLTemplate(t.filename)
		})
	}
	t.templ.Execute(w, nil)
}

func main() {
	flag.Parse()

	http.Handle("/", &templateHandler{filename: "chat.html", debug: *debug})

	if err := http.ListenAndServe(fmt.Sprintf(":%d", *listenPort), nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
