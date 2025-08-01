// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	space := ' '
	hyphen := '-'
	builder := strings.Builder{}
	isBoundary := true
	for _, r := range s {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
			if isBoundary {
				builder.WriteString(strings.ToUpper(string(r)))
				isBoundary = false
			}
		} else if r == space || r == hyphen {
			isBoundary = true
		}

	}
	return builder.String()
}
