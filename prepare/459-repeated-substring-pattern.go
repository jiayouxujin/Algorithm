package main

import "strings"

func repeatedSubstringPattern(s string) bool {
	return strings.Contains((s + s)[1:len(s)*2-1], s)
}
