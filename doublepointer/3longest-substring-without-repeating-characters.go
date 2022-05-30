package main

//给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度
func lengthOfLongestSubstring(s string) int {
	left, window, res := 0, make(map[byte]int), 0

	for right := 0; right < len(s); right++ {
		c := s[right]
		window[c]++

		for left < right && window[c] > 1 {
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
//	res := lengthOfLongestSubstring("abcabcbb")
//	print(res)
//}
