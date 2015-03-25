package angular_play

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var indexTemplate = parseTemplate("indexTemplate", "templates/index.html")

// parseTemplate parses a bunch of templates and assigns a nickname to them.
func parseTemplate(nickname string, filenames ...string) *template.Template {
	t, err := template.New(nickname).Delims("<<", ">>").ParseFiles(filenames...)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

// indexHandler handles the root request URL path.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

// init initializes the module.
func init() {
	http.HandleFunc("/", indexHandler)
}
