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
