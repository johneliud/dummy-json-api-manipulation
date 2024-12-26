package controllers

import (
	"log"
	"net/http"
	"recipes/utils"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		notFoundHandler(w, r)
		log.Println("Page Not Found")
		return
	}

	recipes, err := utils.FetchRecipes()
	if err != nil {
		internalServerErrorHandler(w, r)
		log.Printf("Failed fetching recipes: %v\n", err)
		return
	}

	if len(recipes.Recipes) == 0 {
		internalServerErrorHandler(w, r)
		log.Printf("No available recipes. Fetched %d recipes\n", len(recipes.Recipes))
		return
	}

	tmpl := utils.ParseTemplates()

	if err := tmpl.ExecuteTemplate(w, "index.html", recipes); err != nil {
		internalServerErrorHandler(w, r)
		log.Printf("Error executing template: %v\n", err)
		return
	}
}
