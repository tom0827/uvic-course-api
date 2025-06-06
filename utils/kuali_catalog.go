package utils

import (
	"course-api/constants"
	"course-api/models"
	"encoding/json"
	"fmt"
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

func GetKualiCatalog() ([]models.KualiCourse, error) {
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

	// Set the catalog cache
	catalog.courses = courses
	catalog.expiration = time.Now().Add(cacheDuration)
	return courses, nil
}

func GetKualiCourseInfo(pid string) (*models.KualiCourseInfo, error) {
	var courseInfo models.KualiCourseInfo
	resp, err := http.Get(fmt.Sprintf(constants.InformationUrl, pid))

	if err != nil {
		return nil, fmt.Errorf("failed to fetch course info: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch course info: status %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&courseInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to decode course info: %w", err)
	}

	return &courseInfo, nil
}

func SearchKualiCatalog(courses []models.KualiCourse, search string) []models.KualiCourseSummary {
	var matches []models.KualiCourseSummary
	for _, course := range courses {
		if search == "" || strings.HasPrefix(course.CatalogCourseId, search) {
			matches = append(matches, models.KualiCourseSummary{
				CatalogCourseId: course.CatalogCourseId,
				Pid:             course.Pid,
				Title:           course.Title,
			})
		}
	}
	return matches
}
