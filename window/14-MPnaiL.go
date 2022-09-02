package main

func checkInclusion2(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	left, right, window, freq, count := 0, 0, make(map[byte]int), make(map[byte]int), 0
	for i := 0; i < len(s1); i++ {
		freq[s1[i]]++
	}

	for right < len(s2) {
		c := s2[right]
		right++

		if v, ok := freq[c]; ok {
			window[c]++
			if window[c] == v {
				count++
			}
		}

		for right-left >= len(s1) {
			if count == len(freq) {
				return true
			}
			d := s2[left]
			left++

			if v, ok := window[d]; ok {
				if v == freq[d] {
					count--
				}
				window[d]--
			}
		}
	}
	return false
}
