package main

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	used := make([]bool, len(nums))
	backtrack(nums, []int{}, &used, &res)
	return res
}

func backtrack(nums, tmp []int, used *[]bool, res *[][]int) {
	if len(tmp) == len(nums) {
		c := make([]int, len(nums))
		copy(c, tmp)
		*res = append(*res, c)
		return
	}

	for i := 0; i < len(nums); i++ {
		if (*used)[i] {
			continue
		}
		(*used)[i] = true
		tmp = append(tmp, nums[i])
		backtrack(nums, tmp, used, res)
		tmp = tmp[:len(tmp)-1]
		(*used)[i] = false
	}
}
