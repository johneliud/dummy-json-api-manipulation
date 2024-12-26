package main

import (
	"fmt"
	"net/http"
	"recipes/controllers"
	"recipes/utils"
)

func main() {
	port := utils.GetPort()

	http.HandleFunc("/", controllers.IndexHandler)
	http.HandleFunc("/recipe/", controllers.RecipeDetailsHandler)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Printf("Server started at http://localhost%v\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
