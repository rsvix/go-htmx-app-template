package utils

import (
	"regexp"
	"unicode"
)

// ##########################################
// ###############  Version 1 ###############
// ##########################################

func IsValidPasswordV1(s string) bool {
	if regexp.MustCompile(`\s`).MatchString(s) {
		return false
	}
	if len(s) < 2 {
		return false
	}
	return true
}

// ##########################################
// ###############  Version 2 ###############
// ##########################################

func IsValidPasswordV2(s string) bool {
	containNumber, containUpper, containSpecial, containLower := false, false, false, false
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			containNumber = true
		case unicode.IsUpper(c):
			containUpper = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			containSpecial = true
		case unicode.IsLower(c) || c == ' ':
			containLower = true
		default:
			return false
		}
	}
	if !containNumber || !containUpper || !containSpecial || !containLower {
		return false
	}
	if len(s) < 7 {
		return false
	}
	if regexp.MustCompile(`\s`).MatchString(s) {
		return false
	}
	return true
}
