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

func Mod10CheckDigit(digits string) bool {
	var checksum int

	size := len(digits)
	for i := size - 1; i >= 0; i -= 2 {
		n := digits[i] - '0'
		checksum += int(n)
	}
	for i := size - 2; i >= 0; i -= 2 {
		n := digits[i] - '0'
		n *= 2
		if n > 9 {
			n -= 9
		}
		checksum += int(n)
	}

	return checksum%10 == 0
}
