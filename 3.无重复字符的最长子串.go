/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {

	left,right,length,m:=0,0,0,map[byte]int{}

	for right<len(s){
		c:=s[right]
		right++
		m[c]++

		for m[c]>1{
			remove:=s[left]
			m[remove]--
			left++
		}

		if right-left>length{
			length=right-left
		}
	}
	return length
}
// @lc code=end

