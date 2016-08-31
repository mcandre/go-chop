package chop

import "strings"

const LineEndingCharacters = "\r\n"

func Chop(s string) string {
	return s[:len(s)-1]
}


func Chomp(s string) string {
	return strings.TrimRight(s, LineEndingCharacters)
}

