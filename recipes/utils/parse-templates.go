package utils

import (
	"log"
	"path/filepath"
	"text/template"
)

func ParseTemplates() *template.Template {
	path := filepath.Join("templates", "*.html")

	tmpl, err := template.ParseGlob(path)
	if err != nil {
		log.Printf("Failed parsing template %v: %v\n", path, err)
		return nil
	}
	return tmpl
}
