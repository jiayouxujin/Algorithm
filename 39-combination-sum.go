package main

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	backtrackCombination(candidates, target, 0, 0, &res, []int{})
	return res
}

func backtrackCombination(candidates []int, target, sum, index int, res *[][]int, tmp []int) {
	if sum == target {
		t := make([]int, len(tmp))
		copy(t, tmp)
		*res = append(*res, t)
		return
	}
	if sum > target {
		return
	}
	for i := index; i < len(candidates); i++ {
		tmp = append(tmp, candidates[i])
		backtrackCombination(candidates, target, sum+candidates[i], i, res, tmp)
		tmp = tmp[:len(tmp)-1]
	}
}
