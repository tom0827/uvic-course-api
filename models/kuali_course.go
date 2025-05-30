package models

import "fmt"

type SubjectCode struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Id          string `json:"id"`
	LinkedGroup string `json:"linkedGroup"`
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

// GetCourseByPID returns the first course with the matching PID
func GetCourseByPID(courses []KualiCourse, pid string) (*KualiCourse, error) {
	for i := range courses {
		if courses[i].Pid == pid {
			return &courses[i], nil
		}
	}
	return nil, fmt.Errorf("course with PID %s not found", pid)
}

// GetCourseByDepartmentAndNumber returns the first course matching the department and number
func GetCourseByDepartmentAndNumber(courses []KualiCourse, course string) (*KualiCourse, error) {
	for i := range courses {
		if courses[i].CatalogCourseId == course {
			return &courses[i], nil
		}
	}
	return nil, fmt.Errorf("course %s not found", course)
}
