package main

func canJump(nums []int) bool {
	if nums == nil || len(nums) == 0 {
		return true
	}
	maxEnd := nums[0]
	for i := 1; i < len(nums); i++ {
		if i > maxEnd {
			return false
		}
		maxEnd = max(maxEnd, i+nums[i])
	}
	return true
}
