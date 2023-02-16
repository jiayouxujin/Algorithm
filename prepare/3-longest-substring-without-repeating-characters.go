package main

func lengthOfLongestSubstring(s string) int {
	left, right, res, window := 0, 0, 0, make(map[byte]int)
	for right < len(s) {
		c := s[right]
		right++
		window[c]++
		for left < right && window[c] > 1 {
			d := s[left]
			left++
			window[d]--
		}
		res = max(res, right-left)
	}
	return res
}
