package controllers

import (
	"log"
	"net/http"
	"recipes/models"
	"recipes/utils"
)

func errorHandler(w http.ResponseWriter, errorPage models.ErrorPage) {
	tmpl := utils.ParseTemplates()

	w.WriteHeader(errorPage.StatusCode)

	if err := tmpl.ExecuteTemplate(w, "error-page.html", errorPage); err != nil {
		http.Error(w, "An Unexpected Error Occurred. Try Again Later", http.StatusInternalServerError)
		log.Printf("Error executing template: %v\n", err)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	errorHandler(w, models.ErrorPage{
		StatusCode:   http.StatusNotFound,
		ErrorMessage: "Not Found",
	})
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	errorHandler(w, models.ErrorPage{
		StatusCode:   http.StatusMethodNotAllowed,
		ErrorMessage: "Method Not Allowed",
	})
}

func internalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	errorHandler(w, models.ErrorPage{
		StatusCode:   http.StatusInternalServerError,
		ErrorMessage: "An Unexpected Error Occurred. Try Again Later",
	})
}
