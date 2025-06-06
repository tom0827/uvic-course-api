package utils

import (
	"course-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WriteSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, models.APIResponse{
		Status: "success",
		Data:   data,
	})
}

func WriteError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, models.APIResponse{
		Status:  "error",
		Message: message,
		Data:    nil,
	})
}
