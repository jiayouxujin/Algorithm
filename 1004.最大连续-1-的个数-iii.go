/*
 * @lc app=leetcode.cn id=1004 lang=golang
 *
 * [1004] 最大连续1的个数 III
 */

// @lc code=start
func longestOnes(nums []int, k int) int {
	left, right, res := 0, 0, 0
	for left < len(nums) {
		if right < len(nums) && ((nums[right] == 0 && k > 0) || nums[right] == 1) {
			if nums[right] == 0 {
				k--
			}
			right++
		} else {
			if right==len(nums)||k==0{
				res=max(res,right-left)
			}
			if nums[left]==0{
				k++
			}
			left++
		}
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

