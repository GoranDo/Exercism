// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(greeting string) string {
	greeting = strings.TrimSpace(greeting)

	switch {
	case greeting == "":
		return "Fine. Be that way!"
	case any(greeting, unicode.IsUpper) && !any(greeting, unicode.IsLower) && !strings.ContainsAny(greeting, "?"):
		return "Whoa, chill out!"
	case any(greeting, unicode.IsUpper) && strings.ContainsAny(greeting, "?") && !any(greeting, unicode.IsLower):
		return "Calm down, I know what I'm doing!"
	case greeting[len(greeting)-1] == '?':
		return "Sure."
	default:
		return "Whatever."
	}
}

func any(sentence string, test func(rune) bool) bool {
	for _, letter := range sentence {
		if test(letter) {
			return true
		}
	}
	return false
}
