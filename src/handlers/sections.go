package handlers

import (
	"course-api/src/constants"
	"course-api/src/redis"
	"course-api/src/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func SectionHandler(c *gin.Context) {
	term := strings.ToUpper(c.Query("term"))
	course := strings.ToUpper(c.Query("course"))

	if term == "" || course == "" {
		utils.WriteError(c, "Please provide both term and course")
		return
	}

	cacheKey := fmt.Sprintf("sections:%s:%s", term, course)
	cached, err := redis.Get(cacheKey)
	if err == nil && cached != "" {
		log.Printf("CACHE HIT: Sections for %s %s", term, course)
		var result map[string]interface{}
		if err := json.Unmarshal([]byte(cached), &result); err == nil {
			utils.WriteSuccess(c, result)
			return
		}
	}

	log.Printf("CACHE MISS: Sections for %s %s", term, course)
	subject, number := utils.SplitCourseCode(course)
	cookieLink := fmt.Sprintf(constants.CookieUrl, term)
	dataLink := fmt.Sprintf(constants.SectionsUrl, term, subject, number)

	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}

	_, err = client.Get(cookieLink)
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
		result := map[string]interface{}{
			"sectionsFetchedCount": 0,
			"sections":             []map[string]any{},
		}
		utils.WriteSuccess(c, result)
		data, _ := json.Marshal(result)
		_ = redis.Set(cacheKey, string(data), 24*time.Hour)
		return
	}

	// Cache result
	data, _ := json.Marshal(result)
	_ = redis.Set(cacheKey, string(data), time.Hour)
	utils.WriteSuccess(c, result)
}
