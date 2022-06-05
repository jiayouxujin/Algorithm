package main

func lengthOfLongestSubstring(s string) int {
	left, window, res := 0, make(map[byte]int), 0
	for right := 0; right < len(s); right++ {
		window[s[right]]++

		for window[s[right]] > 1 {
			d := s[left]
			left++
			window[d]--
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
