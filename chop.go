// Package chop provides utility functions for managing trailing characters, such as line feeds and carriage return line feeds.
package chop

import "strings"

// Version is semver.
const Version = "0.0.2"

// LineEndingCharacters lists some common line ending characters.
const LineEndingCharacters = "\r\n"

// Chop removes one trailing character, if any.
func Chop(s string) string {
	if len(s) == 0 {
		return s
	}

	return s[:len(s)-1]
}

// Chomp removes trailing line end characters, if any.
func Chomp(s string) string {
	return strings.TrimRight(s, LineEndingCharacters)
}
