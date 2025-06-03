package handlers

import (
	"course-api/models"
	"course-api/utils"
	"net/http"

	"github.com/gorilla/mux"
)

/*
function CoursesHandler
/api/courses/info/{course-name}
Finds corresponding course name in the Kuali catalog and returns it as a JSON response
*/
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid := vars["pid"]

	courseInfo, err := models.GetKualiCourseInfo(pid)
	if err != nil {
		utils.WriteError(w, "Course not found")
		return
	}

	utils.WriteSuccess(w, courseInfo)
}
