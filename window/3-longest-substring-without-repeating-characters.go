package main

func lengthOfLongestSubstring(s string) int {
	left, window, res := 0, make(map[byte]int), 0
	for right := 0; right < len(s); right++ {
		c := s[right]
		window[c]++

		for window[c] > 1 {
			d := s[left]
			left++
			window[d]--
		}
		if res < right-left+1 {
			res = right - left + 1
		}
	}
	return res
}
