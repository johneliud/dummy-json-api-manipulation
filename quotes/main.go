package main

import (
	"fmt"
	"net/http"

	"quotes/controllers"
	"quotes/utils"
)

func main() {
	port := utils.GetPort()

	http.HandleFunc("/", controllers.IndexHandler)
	http.HandleFunc("/quotes", controllers.QuotesHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	fmt.Printf("Server started at http://localhost%v\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
