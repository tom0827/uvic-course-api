package handlers

import (
	"course-api/constants"
	"course-api/models"
	"course-api/utils"
	"encoding/json"
	"net/http"
)

func CoursesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	course := r.URL.Query().Get("course")
	pid := r.URL.Query().Get("pid")

	validParams := course != "" || pid != ""

	if !validParams {
		http.Error(w, "Invalid parameters", http.StatusBadRequest)
		return
	}

	resp, err := http.Get(constants.CatalogUrl)

	if err != nil {
		http.Error(w, "Failed to fetch courses", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	var courses []models.KualiCourse
	err = json.NewDecoder(resp.Body).Decode(&courses)

	if err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusInternalServerError)
		return
	}

	if pid != "" {
		foundCourse, err := models.GetCourseByPID(courses, pid)
		if err == nil {
			utils.WriteSuccess(w, foundCourse)
			return
		}
	}

	if course != "" {
		foundCourse, err := models.GetCourseByDepartmentAndNumber(courses, course)
		if err == nil {
			utils.WriteSuccess(w, foundCourse)
			return
		}
	}

	utils.WriteError(w, "Course not found")
}
