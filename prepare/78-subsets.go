package main

func subsets(nums []int) [][]int {
	var res [][]int
	backtraceSubSets(nums, []int{}, 0, &res)
	return res
}

func backtraceSubSets(nums, cur []int, start int, res *[][]int) {
	tmp := make([]int, len(cur))
	copy(tmp, cur)
	*res = append(*res, tmp)
	for i := start; i < len(nums); i++ {
		cur = append(cur, nums[i])
		backtraceSubSets(nums, cur, i+1, res)
		cur = cur[:len(cur)-1]
	}
}
