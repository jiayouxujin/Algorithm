package main

func exchange(nums []int) []int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast]%2 != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
	return nums
}
