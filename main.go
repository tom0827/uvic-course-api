package main

import (
	"course-api/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/courses/{pid}/info", handlers.InfoHandler).Methods("GET")
	r.HandleFunc("/api/courses/{course}/sections/{term}", handlers.SectionHandler).Methods("GET")
	r.HandleFunc("/api/courses", handlers.CourseHandler).Methods("GET")

	http.ListenAndServe(":8080", r)
}
