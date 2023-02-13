package main

import "strings"

func reverseWords(s string) string {
	ss := strings.Fields(s)
	reverse151(&ss, 0, len(ss)-1)
	return strings.Join(ss, " ")
}

func reverse151(arr *[]string, left, right int) {
	for left < right {
		(*arr)[left], (*arr)[right] = (*arr)[right], (*arr)[left]
		left++
		right--
	}
}
