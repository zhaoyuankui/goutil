package strutil

import (
	"bytes"
	"strings"
	"unicode/utf8"
)

// ReplaceInvalidRune replaces invalid runes to '*'.
func ReplaceInvalidRune(s string) string {
	if len(s) == 0 {
		return s
	}
	var buf bytes.Buffer
	startIdx := 0
	idx := 0
	for _, v := range s {
		if v != rune(0xfffd) {
			idx += utf8.RuneLen(v)
			continue
		}
		if startIdx < idx {
			buf.WriteString(s[startIdx:idx])
		}
		buf.WriteString("*")
		idx += 1
		startIdx = idx
	}
	if startIdx < idx {
		buf.WriteString(s[startIdx:idx])
	}
	return buf.String()
}

// ReplaceInvalidRune2 replaces invalid runes to '*'.
// Use ReplaceInvalidRune instead, which is often faster.
func ReplaceInvalidRune2(s string) string {
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
