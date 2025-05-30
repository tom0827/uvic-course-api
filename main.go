package main

import (
	"log"
	"net/http"
	"course-api/handlers"
)

func main() {
	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/courses", handlers.CoursesHandler)

	http.Handle("/api/", http.StripPrefix("/api", apiMux))

	log.Println("API running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}