package templates

import (
	"text/template"
)

// ParseTemplate -
func ParseTemplate(file string) *template.Template {

	tmpl := template.Must(template.ParseFiles(file))

	return tmpl
}
