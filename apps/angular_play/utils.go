package angular_play

import "html/template"

// parseTemplate parses a bunch of templates and assigns a nickname to them.
func parseTemplate(nickname string, filenames ...string) *template.Template {
	t := template.New(nickname).Delims("<<", ">>")
	t = template.Must(t.ParseFiles(filenames...))
	return t
}

// parseTemplateString parses a content string as a template.
func parseTemplateString(nickname, content string) *template.Template {
	t := template.New(nickname).Delims("<<", ">>")
	return template.Must(t.Parse(content))
}
