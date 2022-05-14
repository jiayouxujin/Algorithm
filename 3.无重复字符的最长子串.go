/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	window, left, right, res := make(map[byte]int), 0, 0, 0
	for right < len(s) {
		c := s[right]
		right++
		window[c]++

		for window[c] > 1 {
			remove := s[left]
			window[remove]--
			left++
		}

		if right-left > res {
			res = right - left
		}
	}
	return res
}

// @lc code=end

