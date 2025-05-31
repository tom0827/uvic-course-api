package handlers

import (
	"course-api/constants"
	"course-api/models"
	"course-api/utils"
	"encoding/json"
	"net/http"
	"strings"
	"sync"
	"time"
)

type catalogCache struct {
	courses    []models.KualiCourse
	expiration time.Time
	mu         sync.RWMutex
}

var (
	cacheDuration = 30 * time.Minute
	catalog       = &catalogCache{}
)

/*
function CoursesHandler
It expects the following query parameters:
- course: The subject subject + number (e.g., "SENG499").
*/
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	course := strings.ToUpper(r.URL.Query().Get("course"))

	validParams := course != ""

	if !validParams {
		http.Error(w, "Invalid parameters", http.StatusBadRequest)
		return
	}

	var courses []models.KualiCourse
	courses, err := getCachedCourses()

	if err != nil {
		http.Error(w, "Failed to fetch courses", http.StatusInternalServerError)
		return
	}

	foundCourse, err := models.GetCourseByDepartmentAndNumber(courses, course)
	if err == nil {
		utils.WriteSuccess(w, foundCourse)
		return
	}

	utils.WriteError(w, "Course not found")
}

/*
function getCachedCourses
Checks if the course catalog is cached and valid using a double-checked locking pattern
*/
func getCachedCourses() ([]models.KualiCourse, error) {
	catalog.mu.RLock()
	// Check cache with read lock (Allows concurrent reads)
	if time.Now().Before(catalog.expiration) && catalog.courses != nil {
		defer catalog.mu.RUnlock()
		return catalog.courses, nil
	}
	catalog.mu.RUnlock()

	catalog.mu.Lock()
	defer catalog.mu.Unlock() // Defer unlocking until the function exits. Guarantees that the lock is released even if an error occurs.

	// Re-check the cache with write lock (Does not allow concurrent reads)
	if time.Now().Before(catalog.expiration) && catalog.courses != nil {
		return catalog.courses, nil
	}

	// Cache miss, fetch catalog from the URL

	resp, err := http.Get(constants.CatalogUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var courses []models.KualiCourse
	err = json.NewDecoder(resp.Body).Decode(&courses)
	if err != nil {
		return nil, err
	}

	catalog.courses = courses
	catalog.expiration = time.Now().Add(cacheDuration)
	return courses, nil
}
