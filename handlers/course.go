package handlers

import (
	"course-api/models"
	"course-api/utils"
	"net/http"
	"strconv"
	"strings"
)

/*
function CoursesHandler
/api/courses
Returns all courses in catalog or matching a course name
Returns a paginated list of courses
Query parameters:
- search: string, optional, search term to filter courses by name
- page: int, optional, page number for pagination (default is 1)
- pageSize: int, optional, number of courses per page (default is 20)
Maximum page size is 100
*/
func CourseHandler(w http.ResponseWriter, r *http.Request) {
	// Maximum page size limit
	const MaxPageSize = 100
	// Default pagination values
	page := 1
	pageSize := 20

	search := strings.ToUpper(r.URL.Query().Get("search"))
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("pageSize")

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
		pageSize = MaxPageSize // Limit page size to a maximum of 100
	}

	var courses []models.KualiCourse
	courses, err := utils.GetKualiCatalog()

	if err != nil {
		http.Error(w, "Failed to fetch courses", http.StatusInternalServerError)
		return
	}

	matches := models.SearchKualiCatalog(courses, search)
	total := len(matches)

	// Pagination logic
	start := min((page-1)*pageSize, total)
	end := min(start+pageSize, total)

	paginatedMatches := matches[start:end]

	// If no matches found, return an empty slice instead of nil
	if len(paginatedMatches) == 0 {
		paginatedMatches = []models.KualiCourseSummary{}
	}

	utils.WriteSuccess(w, map[string]interface{}{
		"courses":    paginatedMatches,
		"count":      len(paginatedMatches),
		"totalCount": total,
		"page":       page,
		"pageSize":   pageSize,
	})
}
