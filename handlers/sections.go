package handlers

import (
	"course-api/constants"
	"course-api/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"

	"github.com/gin-gonic/gin"
)

// SectionHandler godoc
// @Summary Get course sections
// @Description Get sections for a specific course in a specific term
// @Tags courses
// @Accept json
// @Produce json
// @Param term path string true "Term ID (e.g., 202505)"
// @Param course path string true "Course ID (e.g., SENG499)"
// @Success 200 {object} object "Successful response with sections data or empty array if no sections found"
// @Failure 500 {object} object{error=string} "Error when failing to fetch cookie, sections, read response, or parse JSON"
// @Failure 500 {object} object{error=string} "Error when sections count is invalid"
// @Router /courses/sections/{term}/{course} [get]
func SectionHandler(c *gin.Context) {
	term := c.Param("term")
	course := c.Param("course")

	subject, number := utils.SplitCourseCode(course)
	cookieLink := fmt.Sprintf(constants.CookieUrl, term)
	dataLink := fmt.Sprintf(constants.SectionsUrl, term, subject, number)

	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}

	_, err := client.Get(cookieLink)

	if err != nil {
		utils.WriteError(c, "Failed to fetch cookie")
		return
	}

	resp, err := client.Get(dataLink)

	if err != nil {
		utils.WriteError(c, "Failed to fetch sections")
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		utils.WriteError(c, "Failed to read response body")
		return
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)

	if err != nil {
		utils.WriteError(c, "Failed to decode JSON")
		return
	}

	numOfSections, ok := result["sectionsFetchedCount"].(float64)

	if !ok {
		utils.WriteError(c, "Invalid sections count in response")
		return
	}

	if int(numOfSections) == 0 {
		utils.WriteSuccess(c, []map[string]any{})
		return
	}

	utils.WriteSuccess(c, result)
}
