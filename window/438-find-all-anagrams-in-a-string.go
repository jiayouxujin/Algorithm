package main

func findAnagrams(s string, p string) []int {
	left, right, count, freq, window, res := 0, 0, 0, make(map[byte]int), make(map[byte]int), make([]int, 0)
	for i := 0; i < len(p); i++ {
		freq[p[i]]++
	}
	for right < len(s) {
		c := s[right]
		right++

		if v, ok := freq[c]; ok {
			window[c]++
			if window[c] == v {
				count++
			}
		}

		for right-left >= len(p) {
			if count == len(freq) {
				res = append(res, left)
			}

			d := s[left]
			left++
			if v, ok := freq[d]; ok {
				if window[d] == v {
					count--
				}
				window[d]--
			}
		}
	}
	return res
}
