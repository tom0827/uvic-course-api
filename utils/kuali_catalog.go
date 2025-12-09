package utils

import (
	"course-api/constants"
	"course-api/models"
	"course-api/redis"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const redisCatalogKey = "kuali_catalog"

func GetKualiCatalog() ([]models.KualiCourse, error) {

	// Try Redis cache first
	cached, err := redis.Get(redisCatalogKey)
	if err == nil && cached != "" {
		log.Printf("CACHE HIT: Kuali Catalog")
		var courses []models.KualiCourse
		if err := json.Unmarshal([]byte(cached), &courses); err == nil {
			return courses, nil
		}
	}
	log.Printf("CACHE MISS: Kuali Catalog")
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

	// Set Redis cache
	data, _ := json.Marshal(courses)
	_ = redis.Set(redisCatalogKey, string(data), constants.CacheDuration)
	return courses, nil
}

func GetKualiCourseInfo(pid string, course string) (*models.KualiCourseInfo, error) {
	if pid == "" {
		var courses []models.KualiCourse
		courses, err := GetKualiCatalog()

		if err != nil {
			return nil, fmt.Errorf("failed to fetch Kuali catalog: %w", err)
		}

		matches := SearchKualiCatalog(courses, course)
		if len(matches) == 1 {
			pid = matches[0].Pid
		} else if len(matches) == 0 {
			return nil, fmt.Errorf("no course found with the given course ID")
		} else {
			return nil, fmt.Errorf("multiple courses found with the given course ID, please specify a PID")
		}
	}

	// Redis cache key
	cacheKey := fmt.Sprintf("courseinfo:%s", pid)
	cached, err := redis.Get(cacheKey)
	if err == nil && cached != "" {
		log.Printf("CACHE HIT: Course Info for PID %s", pid)
		var courseInfo models.KualiCourseInfo
		if err := json.Unmarshal([]byte(cached), &courseInfo); err == nil {
			return &courseInfo, nil
		}
	}

	log.Printf("CACHE MISS: Course Info for PID %s", pid)
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

	// Cache result
	data, _ := json.Marshal(courseInfo)
	_ = redis.Set(cacheKey, string(data), constants.CacheDuration)
	return &courseInfo, nil
}

// TODO: Optimize to exit search if we only need one result
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
