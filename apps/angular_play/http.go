package angular_play

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func init() {
	http.HandleFunc('/', indexHandler)
}
