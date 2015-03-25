package angular_play

import (
	"html/template"
	"net/http"
)

type ApplicationConfig struct {
	Title string
}

var indexTemplate = parseTemplate(
	"foobar",
	"templates/base.html",
	"templates/index.html")

var testTemplate = parseTemplateString("foobar", `Example content.`)

// parseTemplate parses a bunch of templates and assigns a nickname to them.
func parseTemplate(nickname string, filenames ...string) *template.Template {
	t := template.New(nickname).Delims("<<", ">>")
	t = template.Must(t.ParseFiles(filenames...))
	return t
}

func parseTemplateString(nickname, content string) *template.Template {
	t := template.New(nickname).Delims("<<", ">>")
	return template.Must(t.Parse(content))
}

// indexHandler handles the root request URL path.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	applicationConfig := ApplicationConfig{
		Title: "MeowApp",
	}

	if err := indexTemplate.ExecuteTemplate(w, "application", applicationConfig); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// init initializes the module.
func init() {
	http.HandleFunc("/", indexHandler)
}
