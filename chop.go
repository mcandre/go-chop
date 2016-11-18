package chop

import "strings"

const Version = "0.0.1"

const LineEndingCharacters = "\r\n"

func Chop(s string) string {
	if len(s) == 0 {
		return s
	}

	return s[:len(s)-1]
}

func Chomp(s string) string {
	return strings.TrimRight(s, LineEndingCharacters)
}
