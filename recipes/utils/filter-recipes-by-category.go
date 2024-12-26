package utils

import (
	"recipes/models"
	"strings"
)

func FilterRecipesByCategory(recipes []models.Recipe, category string) []models.Recipe {
	var filteredRecipes []models.Recipe

	category = strings.ToLower(category)

	for _, recipe := range recipes {
		var isInMealType bool
		for _, mealType := range recipe.MealType {
			if strings.ToLower(mealType) == category {
				isInMealType = true
				break
			}
		}

		var isInTags bool
		for _, tag := range recipe.Tags {
			if strings.ToLower(tag) == category {
				isInTags = true
				break
			}
		}

		if isInMealType || isInTags {
			filteredRecipes = append(filteredRecipes, recipe)
		}
	}
	return filteredRecipes
}
