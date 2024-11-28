package utils

import (
	"regexp"
	"strconv"
)

type BoletoCodeType string

type BoletoType string

const (
	CreditCard         BoletoType = "CREDIT_CARD"
	CityHalls          BoletoType = "CITY_HALLS"
	Sanitation         BoletoType = "SANITATION"
	ElectricityAndGas  BoletoType = "ELECTRICITY_AND_GAS"
	Telecommunications BoletoType = "TELECOMMUNICATIONS"
	GovernmentAgencies BoletoType = "GOVERNMENT_AGENCIES"
	PaymentBooklets    BoletoType = "PAYMENT_BOOKLETS"
	TrafficFines       BoletoType = "TRAFFIC_FINES"
	Bank               BoletoType = "BANK"
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

func CalculateVerificationDigit(block string) string {
	sum := 0
	multiplier := 2

	// Iterate through digits from right to left
	for i := len(block) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(block[i]))
		result := digit * multiplier

		// If result > 9, sum its digits
		if result > 9 {
			result = (result / 10) + (result % 10)
		}

		sum += result

		// Alternate multiplier between 2 and 1
		multiplier = 3 - multiplier
	}

	// Calculate check digit
	remainder := sum % 10
	if remainder == 0 {
		return "0"
	}

	return strconv.Itoa(10 - remainder)
}
