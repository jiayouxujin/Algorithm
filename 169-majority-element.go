package main

func majorityElement(nums []int) int {
	var target, count int
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			target = nums[i]
			count++
		} else if count > 0 {
			if target == nums[i] {
				count++
			} else {
				count--
			}
		}
	}
	return target
}
