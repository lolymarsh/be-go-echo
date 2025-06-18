package util

import (
	"regexp"
	"slices"
	"strings"
	"time"
	"unicode"
)

func StringIsEmpty(text string) bool {
	trimmedText := strings.TrimSpace(text)
	return trimmedText == ""
}

func StringIsNotEmpty(text string) bool {
	trimmedText := strings.TrimSpace(text)
	return trimmedText != ""
}

func StringIsNumber(text string) bool {
	for _, char := range text {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func IsValidHaveStringIn(sortName string, validColumns []string) bool {

	if StringIsEmpty(sortName) {
		return true
	}

	if len(validColumns) == 0 {
		return true
	}

	return slices.Contains(validColumns, sortName)
}

func IsValidRegex(password string) bool {
	pattern := "^[A-Za-z0-9!@#$%^&*()_+-=]+$"
	match, _ := regexp.MatchString(pattern, password)
	return match
}

func IsValidOrder(order string) bool {
	order = strings.ToUpper(strings.TrimSpace(order))
	return order == "ASC" || order == "DESC"
}

func IsValidDateFormat(dateStr, format string) bool {
	_, err := time.Parse(format, dateStr)
	return err == nil
}
