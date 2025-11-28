package handlers

import (
	"course-api/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// CourseInfoHandler godoc
// @Summary Get course information
// @Description Get general information about available courses
// @Tags courses
// @Accept json
// @Produce json
// @Param pid query string false "Program ID (in uppercase)"
// @Param course query string false "Course code (in uppercase)"
// @Success 200 {object} map[string]interface{} "Successful response with course information"
// @Failure 500 {object} object{error=string} "Error when neither pid nor course is provided"
// @Failure 404 {object} object{error=string} "Course not found"
// @Router /courses/info [get]
func CourseInfoHandler(c *gin.Context) {
	pid := c.Query("pid")
	course := strings.ToUpper(c.Query("course"))

	if pid == "" && course == "" {
		utils.WriteError(c, "Please provide either pid or course")
		return
	}

	courseInfo, err := utils.GetKualiCourseInfo(pid, course)

	if err != nil {
		if pid != "" {
			utils.WriteNotFound(c, "Course not found with pid: "+pid)
		} else {
			utils.WriteNotFound(c, "Course not found with course: "+course)
		}
		return
	}

	utils.WriteSuccess(c, courseInfo)
}
