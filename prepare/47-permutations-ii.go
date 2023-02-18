package main

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	used := make([]bool, len(nums))
	backtracePermuteUnique(nums, []int{}, used, &res)
	return res
}

func backtracePermuteUnique(nums, cur []int, used []bool, res *[][]int) {
	if len(cur) == len(nums) {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
			continue
		}
		cur = append(cur, nums[i])
		used[i] = true
		backtracePermuteUnique(nums, cur, used, res)
		cur = cur[:len(cur)-1]
		used[i] = false
	}
}
