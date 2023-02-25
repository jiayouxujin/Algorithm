package main

func sortArray(nums []int) []int {
	stop := true
	for i := len(nums) - 1; i >= 0; i-- {
		stop = true
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				stop = false
			}
		}
		if stop {
			break
		}
	}
	return nums
}
