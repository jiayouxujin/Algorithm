package main

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		tmp := nums[left] + nums[right]
		if tmp == target {
			return []int{nums[left], nums[right]}
		} else if tmp > target {
			right--
		} else {
			left++
		}
	}
	return []int{}
}
