package handlers

import (
	"course-api/models"
	"course-api/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// CourseHandler godoc
// @Summary Search the catalog of UVic courses
// @Description Get a list of all courses, with optional search and pagination
// @Tags courses
// @Accept json
// @Produce json
// @Param search query string false "Search courses by department code prefix (e.g., 'CSC' returns CSC110, CSC111, CSC320, etc.)"
// @Param page query integer false "Page number for pagination (default: 1)" default(1)
// @Param size query integer false "Number of results per page (default: 20)" default(20)
// @Success 200 {array} map[string]interface{}
// @Failure 500 {object} object{error=string} "Course not found"
// @Router /courses [get]
func CourseHandler(c *gin.Context) {
	const MaxPageLimit = 100
	page := 1      //Default
	pageSize := 20 //Default

	search := strings.ToUpper(c.Query("search"))
	pageStr := c.Query("page")
	pageLimitStr := c.Query("limit")

	if pageStr != "" {
		pageTemp, err := strconv.Atoi(pageStr)
		if err == nil && pageTemp > 0 {
			page = pageTemp
		}
	}

	if pageLimitStr != "" {
		pageLimitTemp, err := strconv.Atoi(pageLimitStr)
		if err == nil && pageLimitTemp > 0 {
			pageSize = pageLimitTemp
		}
	}

	if pageSize > MaxPageLimit {
		pageSize = MaxPageLimit
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
