package main

import (
	"html/template"
	"path/filepath"
)

// We use these delimiters because client libraries like AngularJS use the
// default template delimiters too.  As AngularJS allows the development of
// reusable components, it is incredibly hard to change the template delimiters
// for client-side templates.
const (
	leftDelim  = "{{{"
	rightDelim = "}}}"
)

// parse parses a sequence of related templates while reading them from
// the specified template path.
func parse(templatePath string, filenames ...string) *template.Template {
	paths := []string{}
	for _, filename := range filenames {
		paths = append(paths, filepath.Join(templatePath, filename))
	}
	return template.Must(template.ParseFiles(paths...)).Delims(leftDelim, rightDelim)
}
