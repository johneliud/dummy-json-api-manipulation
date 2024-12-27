package controllers

import (
	"log"
	"net/http"
	"recipes/models"
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

	data := struct {
		Recipes  []models.Recipe
		Featured models.Recipe
	}{
		Recipes:  recipes.Recipes,
		Featured: recipes.Recipes[0], // Default featured
	}

	tmpl := utils.ParseTemplates()

	if err := tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
		internalServerErrorHandler(w, r)
		log.Printf("Error executing template: %v\n", err)
		return
	}
}
