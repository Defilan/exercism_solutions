package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram checks if string is an isogram
func IsIsogram(word string) bool {
	charexists := map[rune]bool{}
	for _, char := range strings.ToLower(word) {
		if unicode.IsLetter(char) && charexists[char] {
			return false
		}
		charexists[char] = true
	}
	return true
}