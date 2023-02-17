package main

func jump(nums []int) int {
	n := len(nums)
	step, end, maxEnd := 0, 0, 0
	for i := 0; i < n-1; i++ {
		maxEnd = max(maxEnd, i+nums[i])
		if i == end {
			step++
			end = maxEnd
		}
	}
	return step
}
