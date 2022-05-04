/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	if s1 == "" {
		return true
	}
	if s2 == "" {
		return false
	}

	need, window := map[byte]int{}, map[byte]int{}

	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}

	left, right, count := 0, 0, 0
	for right < len(s2) {
		c := s2[right]
		right++

		if _, ok := need[c]; ok {
			window[c]++
			if need[c] == window[c] {
				count++
			}
		}

		for right-left >= len(s1) {
			if count == len(need) {
				return true
			}

			d := s2[left]
			left++

			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					count--
				}
				window[d]--
			}
		}
	}
	return false
}

// @lc code=end

