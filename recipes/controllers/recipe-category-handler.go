package controllers

import (
	"log"
	"net/http"
	"recipes/models"
	"recipes/utils"
	"strings"
)

func RecipeCategoryHandler(w http.ResponseWriter, r *http.Request) {
	category := strings.TrimPrefix(r.URL.Path, "/")

	recipes, err := utils.FetchRecipes()
	if err != nil {
		internalServerErrorHandler(w, r)
		log.Printf("Failed fetching recipes: %v\n", err)
		return
	}

	filteredRecipes := utils.FilterRecipesByCategory(recipes.Recipes, category)

	data := struct {
		Recipes  []models.Recipe
		Category string
	}{
		Recipes:  filteredRecipes,
		Category: strings.ToUpper(string(category[0])) + category[1:],
	}

	tmpl := utils.ParseTemplates()

	if err := tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
		internalServerErrorHandler(w, r)
		log.Printf("Error executing template: %v\n", err)
		return
	}
}
