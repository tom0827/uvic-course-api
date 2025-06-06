package models

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
