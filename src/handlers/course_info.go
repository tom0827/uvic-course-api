package handlers

import (
	"course-api/src/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

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
