package main

func findAnagrams(s string, p string) []int {
	var res []int
	left, right, count, window, need := 0, 0, 0, make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}
	for left <= right && right < len(s) {
		d := s[right]
		right++
		if v, ok := need[d]; ok {
			window[d]++
			if window[d] == v {
				count++
			}
		}
		for right-left >= len(p) {
			if count == len(need) {
				res = append(res, left)
			}
			d := s[left]
			left++
			if v, ok := need[d]; ok {
				if window[d] == v {
					count--
				}
				window[d]--
			}
		}
	}
	return res
}
