package models

type Recipes struct {
	Recipes []Recipe `json:"recipes"`
}

type Recipe struct {
	ID                 int      `json:"id"`
	Name               string   `json:"name"`
	Ingredients        []string `json:"ingredients"`
	Instructions       []string `json:"instructions"`
	PrepTimeMinutes    int      `json:"prepTimeMinutes"`
	CookTimeMinutes    int      `json:"cookTimeMinutes"`
	Servings           int      `json:"servings"`
	Difficulty         string   `json:"difficulty"`
	Cuisine            string   `json:"cuisine"`
	CaloriesPerServing int      `json:"caloriesPerServing"`
	Tags               []string `json:"tags"`
	UserID             int      `json:"userId"`
	Image              string   `json:"image"`
	Rating             float64  `json:"rating"`
	ReviewCount        int      `json:"reviewCount"`
	MealType           []string `json:"mealType"`
}

type ErrorPage struct {
	StatusCode   int
	ErrorMessage string
}
