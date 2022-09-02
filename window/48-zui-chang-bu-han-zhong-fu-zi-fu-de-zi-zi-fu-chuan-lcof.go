package main

func lengthOfLongestSubstring2(s string) int {
	left, right, window, res := 0, 0, make(map[byte]int), 0
	for right < len(s) {
		c := s[right]
		right++

		window[c]++

		for window[c] > 1 {
			d := s[left]
			left++

			window[d]--
		}

		if right-left > res {
			res = right - left
		}
	}
	return res
}
