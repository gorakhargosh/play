package main

import (
	"flag"
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

// The host and port to which the HTTP server will bind.
var listenAddr = flag.String("addr", ":8080", "listen on <host:port>")

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

// Entry-point function.
func main() {
	flag.Parse()

	r := newRoom()

	http.Handle("/", &templateHandler{filename: "chat.html", debug: *debug})
	http.Handle("/room", r)

	go r.run()

	log.Println("Started HTTP server on: ", *listenAddr)
	if err := http.ListenAndServe(*listenAddr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
