/*
 * @lc app=leetcode.cn id=1438 lang=golang
 *
 * [1438] 绝对差不超过限制的最长连续子数组
 */

// @lc code=start
func longestSubarray(nums []int, limit int) int {
	minStack, maxStack := []int{}, []int{}
	left, res := 0, 0
	for right, v := range nums {
		for len(minStack) > 0 && nums[minStack[0]] > v {
			minStack = minStack[1:]
		}
		minStack = append(minStack, right)

		for len(maxStack) > 0 && nums[maxStack[0]] < v {
			maxStack = maxStack[1:]
		}
		maxStack = append(maxStack, right)

		for len(maxStack) > 0 && len(minStack) > 0 && nums[maxStack[0]]-nums[minStack[0]] > limit {
			if maxStack[0] == left {
				maxStack = maxStack[1:]
			}
			if minStack[0] == left {
				minStack = minStack[1:]
			}
			left++
		}

		if right-left+1 > res {
			res = right - left + 1
		}
	}
	return res
}

// @lc code=end

