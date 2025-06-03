package models

import (
	"course-api/constants"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Credits struct {
	Min string `json:"min"`
	Max string `json:"max"`
}

type CreditsInfo struct {
	Credits Credits `json:"credits"`
	Value   string  `json:"value"`
	Chosen  string  `json:"chosen"`
}

type SubjectCode struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Id          string `json:"id"`
	LinkedGroup string `json:"linkedGroup"`
}

type KualiCourseInfo struct {
	Description        string      `json:"description"`
	Pid                string      `json:"pid"`
	Title              string      `json:"title"`
	Recommendations    string      `json:"recommendations"`
	CatalogCourseId    string      `json:"__catalogCourseId"`
	CreditsInfo        CreditsInfo `json:"credits"`
	PreAndCorequisites string      `json:"preAndCorequisites"`
	SubjectCode        SubjectCode `json:"subjectCode"`
	HoursCatalogText   string      `json:"hours"`
}

type KualiCourse struct {
	CatalogCourseId       string      `json:"__catalogCourseId"`
	PassedCatalogQuery    bool        `json:"__passedCatalogQuery"`
	DateStart             string      `json:"dateStart"`
	Pid                   string      `json:"pid"`
	Title                 string      `json:"title"`
	SubjectCode           SubjectCode `json:"subjectCode"`
	CatalogActivationDate string      `json:"catalogActivationDate"`
	Score                 int         `json:"score"`
}

type KualiCourseSummary struct {
	CatalogCourseId string `json:"catalogCourseId"`
	Pid             string `json:"pid"`
	Title           string `json:"title"`
}

func GetKualiCourseInfo(pid string) (*KualiCourseInfo, error) {
	var courseInfo KualiCourseInfo
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

func SearchKualiCatalog(courses []KualiCourse, search string) []KualiCourseSummary {
	var matches []KualiCourseSummary
	for _, course := range courses {
		if search == "" || strings.HasPrefix(course.CatalogCourseId, search) {
			matches = append(matches, KualiCourseSummary{
				CatalogCourseId: course.CatalogCourseId,
				Pid:             course.Pid,
				Title:           course.Title,
			})
		}
	}
	return matches
}
