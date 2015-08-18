package main

// parse parses templates given the templatePath and some filenames.
import (
	"html/template"
	"path/filepath"
)

const (
	leftDelim  = "{{{"
	rightDelim = "}}}"
)

func parse(templatePath string, filenames ...string) *template.Template {
	paths := []string{}
	for _, filename := range filenames {
		paths = append(paths, filepath.Join(templatePath, filename))
	}
	return template.Must(
		template.ParseFiles(paths...)).Delims(leftDelim, rightDelim)
}
