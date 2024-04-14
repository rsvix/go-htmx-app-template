package utils

import (
	"unicode"
)

// var IsAlpha = regexp.MustCompile(`^[A-Za-z]+$`).MatchString

func IsValidName(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
