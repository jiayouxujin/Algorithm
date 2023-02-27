package main

func towStackSort(nums []int) []int {
	tmp := make([]int, 0)
	for len(nums) > 0 {
		v := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		for len(tmp) > 0 && tmp[len(tmp)-1] > v {
			t := tmp[len(tmp)-1]
			tmp = tmp[:len(tmp)-1]
			nums = append(nums, t)
		}
		tmp = append(tmp, v)
	}
	return tmp
}
