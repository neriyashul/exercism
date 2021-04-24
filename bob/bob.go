// Package bob implements Bob's responses for remarks.
package bob

import (
	"strings"
)

// Hey returns Bob's response according to a remark.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	isQuestion := strings.HasSuffix(remark, "?")
	isYelling := (remark == strings.ToUpper(remark) && hasLetter(remark))

	switch {
	case isYelling && isQuestion:
		return "Calm down, I know what I'm doing!"
	case strings.HasSuffix(remark, "?"):
		return "Sure."
	case isYelling:
		return "Whoa, chill out!"
	case remark == "":
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}

func hasLetter(remark string) bool {
	for _, c := range remark {
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
			return true
		}
	}
	return false
}
