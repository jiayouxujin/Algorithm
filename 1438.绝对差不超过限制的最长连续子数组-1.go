/*
 * @lc app=leetcode.cn id=1438 lang=golang
 *
 * [1438] 绝对差不超过限制的最长连续子数组
 */

// @lc code=start
func longestSubarray(nums []int, limit int) int {
	minStack, maxStack, left, res := []int{}, []int{}, 0, 0
	for right, num := range nums {
		for len(minStack) > 0 && nums[minStack[len(minStack)-1]] > num {
			minStack = minStack[:len(minStack)-1]
		}
		minStack = append(minStack, right)
		for len(maxStack) > 0 && nums[maxStack[len(maxStack)-1]] < num {
			maxStack = maxStack[:len(maxStack)-1]
		}
		maxStack = append(maxStack, right)
		if len(minStack) > 0 && len(maxStack) > 0 && nums[maxStack[0]]-nums[minStack[0]] > limit {
			if left == minStack[0] {
				minStack = minStack[1:]
			}
			if left == maxStack[0] {
				maxStack = maxStack[1:]
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

