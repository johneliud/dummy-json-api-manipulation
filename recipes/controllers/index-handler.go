package controllers

import (
	"log"
	"net/http"
	"recipes/utils"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := utils.ParseTemplates()

	if err := tmpl.ExecuteTemplate(w, "index.html", nil); err != nil {
		internalServerErrorHandler(w, r)
		log.Printf("Error executing template: %v\n", err)
	}
}
