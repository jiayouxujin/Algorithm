package main

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	backtrackSubsets(nums, 0, []int{}, &res)
	return res
}

func backtrackSubsets(nums []int, start int, tmp []int, res *[][]int) {
	t := make([]int, len(tmp))
	copy(t, tmp)
	*res = append(*res, t)
	for i := start; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
		backtrackSubsets(nums, i+1, tmp, res)
		tmp = tmp[:len(tmp)-1]
	}
}
