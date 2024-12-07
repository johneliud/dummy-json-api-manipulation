package controllers

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"quotes/utils"
)

// IndexHandler serves the index page by populating fetched data.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Println("Not found.")
		http.NotFound(w, r)
		return
	}

	quotes, err := utils.FetchQuotes()
	if err != nil {
		log.Printf("Error fetching quotes: %v\n", err)
		http.Error(w, fmt.Sprintf("Error fetching quotes: %v", err), http.StatusInternalServerError)
		return
	}

	if len(quotes.Quotes) == 0 {
		log.Printf("No quotes available. Fetched %d quotes.\n", len(quotes.Quotes))
		http.Error(w, "No quotes available", http.StatusInternalServerError)
		return
	}

	path := filepath.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		log.Printf("Error parsing template: %v\n", err)
		http.Error(w, fmt.Sprintf("Error parsing template: %v", err), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, quotes); err != nil {
		log.Printf("Error executing template: %v\n", err)
		http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
	}
}
