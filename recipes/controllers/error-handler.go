package controllers

import (
	"log"
	"net/http"
	"path/filepath"
	"recipes/models"
	"text/template"
)

func errorHandler(w http.ResponseWriter, errorPage models.ErrorPage) {
	path := filepath.Join("templates", "error-page.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, "An Unexpected Error Occurred. Try Again Later", http.StatusInternalServerError)
		log.Printf("Error parsing template %s: %v\n", path, err)
		return
	}
	w.WriteHeader(errorPage.StatusCode)

	if err := tmpl.Execute(w, errorPage); err != nil {
		http.Error(w, "An Unexpected Error Occurred. Try Again Later", http.StatusInternalServerError)
		log.Printf("Error executing template: %v\n", err)
	}
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	errorHandler(w, models.ErrorPage{
		StatusCode:   http.StatusNotFound,
		ErrorMessage: "Not Found",
	})
}

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	errorHandler(w, models.ErrorPage{
		StatusCode:   http.StatusMethodNotAllowed,
		ErrorMessage: "Method Not Allowed",
	})
}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	errorHandler(w, models.ErrorPage{
		StatusCode:   http.StatusInternalServerError,
		ErrorMessage: "An Unexpected Error Occurred. Try Again Later",
	})
}
