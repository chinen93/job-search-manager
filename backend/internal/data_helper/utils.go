package datahelper

import (
	"regexp"
)

var allowedChars = regexp.MustCompile(`[^a-zA-Z0-9_]`)

func SanitizeString(input string) string {

	sanitized := allowedChars.ReplaceAllString(input, "")

	return sanitized
}

// Helper function to compare two string slices
func Equal(first, second []string) bool {
	if len(first) != len(second) {
		return false
	}
	for index, value := range first {
		if value != second[index] {
			return false
		}
	}
	return true
}
