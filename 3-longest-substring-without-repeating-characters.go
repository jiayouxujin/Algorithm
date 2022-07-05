package main

func lengthOfLongestSubstring(s string) int {
	left, res, window := 0, 0, make(map[byte]int)
	for right := 0; right < len(s); right++ {
		window[s[right]]++

		for window[s[right]] > 1 {
			d := s[left]
			window[d]--
			left++
		}
		if right-left+1 > res {
			res = right - left + 1
		}
	}
	return res
}

//func main() {
//	s := "abcabcbb"
//	fmt.Printf("%v \n", lengthOfLongestSubstring(s))
//}
