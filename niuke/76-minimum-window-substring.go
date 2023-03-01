package main

func minWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}
	left, right, window, need, res, count, l := 0, 0, make(map[byte]int), make(map[byte]int), "", 0, len(s)+1
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	for right < len(s) {
		c := s[right]
		right++
		if need[c] > 0 {
			window[c]++
			if need[c] == window[c] {
				count++
			}
		}

		for count == len(need) {
			if l > right-left {
				l = right - left
				res = s[left:right]
			}
			d := s[left]
			left++
			if need[d] > 0 {
				if window[d] == need[d] {
					count--
				}
				window[d]--
			}
		}
	}
	return res
}
