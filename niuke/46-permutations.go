package main

func permute(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return [][]int{}
	}
	res, used := make([][]int, 0), make([]bool, len(nums))
	permuteSolve(nums, []int{}, &res, used)
	return res
}

func permuteSolve(nums, tmp []int, res *[][]int, used []bool) {
	if len(tmp) == len(nums) {
		t := make([]int, len(tmp))
		copy(t, tmp)
		*res = append(*res, t)
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] == false {
			used[i] = true
			tmp = append(tmp, nums[i])
			permuteSolve(nums, tmp, res, used)
			used[i] = false
			tmp = tmp[:len(tmp)-1]
		}
	}
}
