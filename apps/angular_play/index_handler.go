package angular_play

import "net/http"

var indexTemplate = parseTemplate(
	"index",
	"templates/base.html",
	"templates/index.html")

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
