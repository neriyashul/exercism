// Package acronym creates an abbreviation of words.
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate returns an abbreviation of s.
func Abbreviate(s string) string {
	var abbreviation []rune
	for _, word := range strings.FieldsFunc(s, isWordEnd) {
		r := rune(word[0])
		abbreviation = append(abbreviation, unicode.ToUpper(r))
	}

	return string(abbreviation)
}

func isWordEnd(c rune) bool {
	return !unicode.IsLetter(c) && c != '\''
}
