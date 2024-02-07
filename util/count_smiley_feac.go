package util

import "regexp"

func CountSmileyFace(input []string) int {
	pattern := "^[:;][-~]*[)D]$"
	re := regexp.MustCompile(pattern)
	count := 0
	for _, face := range input {
		if re.MatchString(face) {
			count += 1
		}
	}

	return count
}
