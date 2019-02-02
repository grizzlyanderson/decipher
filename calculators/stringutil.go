package calculators

import (
	"strings"
	"unicode"
)

//Normalize string by removing non-letters, insure all upper case
func Normalize(text string) string {
	var result strings.Builder

	for _, b := range text {
		r := rune(b)
		if unicode.IsLetter(r) {
			result.WriteRune(unicode.ToUpper(r))
		}
	}

	return result.String()
}
