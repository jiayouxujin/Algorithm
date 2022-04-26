/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	needCheck := map[byte]int{}
	for i := 0; i < len(t); i++ {
		needCheck[t[i]]++
	}

	window := map[byte]int{}
	left, right, count, start, minLen, ans := 0, 0, 0, -1, len(s)+1, ""
	for right < len(s) {
		c := s[right]
		right++

		if _, ok := needCheck[c]; ok {
			window[c]++
			if window[c] == needCheck[c] {
				count++
			}
		}

		for count == len(needCheck) {
			if right-left < minLen {
				minLen = right - left
				start = left
			}
			r := s[left]
			left++

			if _, ok := needCheck[r]; ok {
				if window[r] == needCheck[r] {
					count--
				}
				window[r]--
			}
		}
	}
	if start != -1 {
		ans = s[start : start+minLen]
	}
	return ans
}

// @lc code=end
