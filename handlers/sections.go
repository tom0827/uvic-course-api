package handlers

import (
	"course-api/constants"
	"course-api/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
)

/*
function SectionHandler
It expects the following query parameters:
- subject: The subject code of the course (e.g., "SENG").
- number: The course number (e.g., "499").
- year: The year of the course (e.g., "2025").
- term: The term of the course (e.g., "05" for Summer).
*/
func SectionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	subject := r.URL.Query().Get("subject")
	number := r.URL.Query().Get("number")
	year := r.URL.Query().Get("year")
	term := r.URL.Query().Get("term")

	validParams := subject != "" && number != "" && year != "" && term != ""

	if !validParams {
		http.Error(w, "Invalid parameters", http.StatusBadRequest)
		return
	}

	yearAndTerm := year + term

	cookieLink := fmt.Sprintf(constants.CookieUrl, yearAndTerm)
	dataLink := fmt.Sprintf(constants.SectionsUrl, yearAndTerm, subject, number)

	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}

	_, err := client.Get(cookieLink)
	if err != nil {
		utils.WriteError(w, "Failed to fetch cookie")
		return
	}

	resp, err := client.Get(dataLink)
	if err != nil {
		utils.WriteError(w, "Failed to fetch sections")
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		utils.WriteError(w, "Failed to read response body")
		return
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		utils.WriteError(w, "Failed to decode JSON")
		return
	}

	utils.WriteSuccess(w, result)
}
