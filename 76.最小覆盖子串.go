/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}

	need, window := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left, right, ans, count, start, minLen := 0, 0, "", 0, -1, len(s)+1
	for right < len(s) {
		c := s[right]
		right++

		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				count++
			}
		}

		for count == len(need) {
			if minLen > right-left {
				minLen = right - left
				start = left
			}

			d := s[left]
			left++

			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					count--
				}
				window[d]--
			}
		}
	}
	if start != -1 {
		ans = s[start : start+minLen]
	}
	return ans
}

// @lc code=end
