package utils

import (
	"regexp"
)

// Substr returns the portion of string specified by the start and length parameters.
func Substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}

// OnlyNumbers returns only numeric characters
func OnlyNumbers(str string) string {
	re := regexp.MustCompile("[^0-9]+")
	return re.ReplaceAllString(str, "")
}
