package main

func sortColors(nums []int) {
	left, right := 0, len(nums)-1
	for i := 0; i <= right; {
		if nums[i] == 0 {
			nums[i], nums[left] = nums[left], nums[i]
			i++
			left++
		} else if nums[i] == 2 {
			nums[i], nums[right] = nums[right], nums[i]
			right--
		} else {
			i++
		}
	}
}
