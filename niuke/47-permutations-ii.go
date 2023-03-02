package main

import "sort"

func permuteUnique(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return [][]int{}
	}
	sort.Ints(nums)
	res, used, tmp := make([][]int, 0), make([]bool, len(nums)), make([]int, 0)
	backtracePermuteUnique(nums, tmp, used, &res)
	return res
}

func backtracePermuteUnique(nums, tmp []int, used []bool, res *[][]int) {
	if len(tmp) == len(nums) {
		t := make([]int, len(tmp))
		copy(t, tmp)
		*res = append(*res, t)
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] || (i-1 >= 0 && nums[i] == nums[i-1] && !used[i-1]) {
			continue
		}
		used[i] = true
		tmp = append(tmp, nums[i])
		backtracePermuteUnique(nums, tmp, used, res)
		used[i] = false
		tmp = tmp[:len(tmp)-1]
	}
}
