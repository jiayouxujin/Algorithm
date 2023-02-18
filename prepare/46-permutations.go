package main

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	used := make([]bool, len(nums))
	backtracePermute(nums, []int{}, used, &res)
	return res
}

func backtracePermute(nums, cur []int, used []bool, res *[][]int) {
	if len(nums) == len(cur) {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		used[i] = true
		cur = append(cur, nums[i])
		backtracePermute(nums, cur, used, res)
		used[i] = false
		cur = cur[:len(cur)-1]
	}
}
