package strutil

import (
	"strings"
)

// ReplaceInvalidRune replaces invalid runes to '*'.
func ReplaceInvalidRune(s string) string {
	if len(s) == 0 {
		return s
	}
	var builder strings.Builder
	var v rune
	for _, v = range s {
		if v != rune(0xfffd) {
			builder.WriteString(string(v))
			continue
		}
		builder.WriteString("*")
	}
	return builder.String()
}
