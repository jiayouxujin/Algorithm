package main

func maxProduct(nums []int) int {
	n := len(nums)
	maxVal, minVal, res := nums[0], nums[0], nums[0]
	for i := 1; i < n; i++ {
		if nums[i] < 0 {
			maxVal, minVal = minVal, maxVal
		}
		maxVal = max(nums[i], nums[i]*maxVal)
		minVal = min(nums[i], nums[i]*minVal)
		res = max(res, maxVal)
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
