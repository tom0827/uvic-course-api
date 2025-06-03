package utils

import "regexp"

func SplitCourseCode(code string) (subject, number string) {
	re := regexp.MustCompile(`^([A-Za-z]+)(\d+[A-Za-z]*)$`)
	matches := re.FindStringSubmatch(code)
	if len(matches) == 3 {
		return matches[1], matches[2]
	}
	return "", ""
}
