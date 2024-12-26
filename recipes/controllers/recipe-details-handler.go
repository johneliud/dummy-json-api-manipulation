package controllers

import (
	"log"
	"net/http"
	"recipes/utils"
	"strconv"
	"strings"
)

func RecipeDetailsHandler(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) != 3 || pathParts[1] != "recipe" {
		notFoundHandler(w, r)
		return
	}

	recipeID, err := strconv.Atoi(string(pathParts[2][0]))
	if err != nil {
		internalServerErrorHandler(w, r)
		log.Printf("Failed converting %v to int: %v\n", pathParts[2][0], err)
		return
	}

	recipe, err := utils.FetchRecipeByID(recipeID)
	if err != nil {
		internalServerErrorHandler(w, r)
		log.Printf("Failed fetching recipe details: %v\n", err)
		return
	}

	tmpl := utils.ParseTemplates()

	if err := tmpl.ExecuteTemplate(w, "recipe-details.html", recipe); err != nil {
		internalServerErrorHandler(w, r)
		log.Printf("Error executing template: %v\n", err)
		return
	}
}
