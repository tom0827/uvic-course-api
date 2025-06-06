package handlers

import (
	"course-api/models"
	"course-api/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CourseHandler(c *gin.Context) {
	const MaxPageSize = 100
	page := 1      //Default
	pageSize := 20 //Default

	search := strings.ToUpper(c.Query("search"))
	pageStr := c.Query("search")
	pageSizeStr := c.Query("search")

	if pageStr != "" {
		pageTemp, err := strconv.Atoi(pageStr)
		if err == nil && pageTemp > 0 {
			page = pageTemp
		}
	}

	if pageSizeStr != "" {
		pageSizeTemp, err := strconv.Atoi(pageSizeStr)
		if err == nil && pageSizeTemp > 0 {
			pageSize = pageSizeTemp
		}
	}

	if pageSize > MaxPageSize {
		pageSize = MaxPageSize
	}

	var courses []models.KualiCourse
	courses, err := utils.GetKualiCatalog()

	if err != nil {
		utils.WriteError(c, "Failed to fetch courses from Kuali catalog")
		return
	}

	matches := utils.SearchKualiCatalog(courses, search)
	total := len(matches)

	// Pagination logic
	start := min((page-1)*pageSize, total)
	end := min(start+pageSize, total)

	paginatedMatches := matches[start:end]

	// If no matches found, return an empty slice instead of nil
	if len(paginatedMatches) == 0 {
		paginatedMatches = []models.KualiCourseSummary{}
	}

	utils.WriteSuccess(c, map[string]interface{}{
		"courses":    paginatedMatches,
		"count":      len(paginatedMatches),
		"totalCount": total,
		"page":       page,
		"pageSize":   pageSize,
	})
}
