package handlers

import (
	"course-api/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func InfoHandler(c *gin.Context) {
	pid := strings.ToUpper(c.Query("pid"))
	course := strings.ToUpper(c.Query("course"))

	if pid == "" && course == "" {
		c.JSON(400, gin.H{"error": "Either pid or course must be provided"})
		return
	}

	courseInfo, err := utils.GetKualiCourseInfo(pid, course)
	if err != nil {
		utils.WriteError(c, "Course not found")
		return
	}

	utils.WriteSuccess(c, courseInfo)
}
