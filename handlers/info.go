package handlers

import (
	"course-api/utils"

	"github.com/gin-gonic/gin"
)

func InfoHandler(c *gin.Context) {
	pid := c.Param("pid")

	courseInfo, err := utils.GetKualiCourseInfo(pid)
	if err != nil {
		utils.WriteError(c, "Course not found")
		return
	}

	utils.WriteSuccess(c, courseInfo)
}
