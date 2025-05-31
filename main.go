package main

import (
	"course-api/handlers"
	"log"
	"net/http"
)

func main() {
	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/info", handlers.InfoHandler)
	apiMux.HandleFunc("/sections", handlers.SectionHandler)

	http.Handle("/api/", http.StripPrefix("/api", apiMux))

	log.Println("API running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
