package main

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	backtrace(candidates, []int{}, 0, 0, target, &res)
	return res
}

func backtrace(candidates, cur []int, sum, index, target int, res *[][]int) {
	if sum > target {
		return
	}
	if sum == target {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}
	for i := index; i < len(candidates); i++ {
		cur = append(cur, candidates[i])
		backtrace(candidates, cur, sum+candidates[i], i, target, res)
		cur = cur[:len(cur)-1]
	}
}
