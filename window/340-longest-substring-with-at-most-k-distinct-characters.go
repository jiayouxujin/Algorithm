package main

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	left, window, res := 0, make(map[byte]int), 0
	for right := 0; right < len(s); right++ {
		c := s[right]
		window[c]++

		for len(window) > k {
			d := s[left]
			left++

			if window[d] == 1 {
				delete(window, d)
			} else {
				window[d]--
			}
		}

		if res < right-left+1 {
			res = right - left + 1
		}
	}
	return res
}

//func main() {
//	s, k := "aa", 1
//	fmt.Printf("%v \n", lengthOfLongestSubstringKDistinct(s, k))
//}
