package main

import (
	"fmt"
	"net/http"

	"quotes/controllers"
)

func main() {
	http.HandleFunc("/", controllers.IndexHandler)
	http.HandleFunc("/quotes", controllers.QuotesHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	fmt.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
