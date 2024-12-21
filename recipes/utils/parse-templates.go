package utils

import (
	"path/filepath"

	"text/template"
)

// ParseTemplate uses a FuncMap to register RenderRatingStars to be called from any of the templates.
func ParseTemplates() *template.Template {
	funcMap := template.FuncMap{
		"renderRatingStars": RenderRatingStars,
	}

	path := filepath.Join("templates", "*.html")
	tmpl := template.Must(
		template.New("").Funcs(funcMap).ParseGlob(path),
	)
	return tmpl
}
