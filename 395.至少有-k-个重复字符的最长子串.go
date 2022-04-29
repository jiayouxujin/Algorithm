/*
 * @lc app=leetcode.cn id=395 lang=golang
 *
 * [395] 至少有 K 个重复字符的最长子串
 */

// @lc code=start
func longestSubstring(s string, k int) int {
	if s == "" {
		return 0
	}
	res, split, freq := 0, byte(0), [26]int{}
	for _, c := range s {
		freq[c-'a']++
	}

	for i, c := range freq[:] {
		if c > 0 && c < k {
			split = 'a' + byte(i)
			break
		}
	}

	if split == 0 {
		return len(s)
	}

	for _, subStr := range strings.Split(s, string(split)) {
		res = max(res, longestSubstring(subStr, k))
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

