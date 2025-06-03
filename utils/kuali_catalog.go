package utils

import (
	"course-api/constants"
	"course-api/models"
	"encoding/json"
	"net/http"
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
function getKualiCatalog
Checks if the course catalog is cached and valid using a double-checked locking pattern
*/
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
