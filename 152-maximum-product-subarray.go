package main

func maxProduct(nums []int) int {
	minimum, maximum, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maximum, minimum = minimum, maximum
		}
		maximum = max(nums[i], nums[i]*maximum)
		minimum = min(nums[i], nums[i]*minimum)
		res = max(maximum, res)
	}
	return res
}
