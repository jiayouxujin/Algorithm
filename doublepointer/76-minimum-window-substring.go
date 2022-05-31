package main

import "fmt"

func minWindow(s string, t string) string {
	if s == "" {
		return ""
	}
	if t == "" {
		return s
	}

	need, window := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left, right, count, begin, length := 0, 0, 0, -1, len(s)
	for right < len(s) {
		c := s[right]
		right++
		window[c]++

		if v, ok := need[c]; ok {
			if v == window[c] {
				count++
			}
		}

		for count == len(need) {
			if right-left <= length {
				begin = left
				length = right - left
			}

			d := s[left]
			left++

			if v, ok := need[d]; ok {
				if v == window[d] {
					count--
				}
			}
			window[d]--
		}
	}
	if begin == -1 {
		return ""
	}
	return s[begin : begin+length]
}

func main() {
	s, t := "ADOBECODEBANC", "ABC"
	fmt.Printf("%v \n", minWindow(s, t))
}
