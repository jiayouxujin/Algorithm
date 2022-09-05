package main

//循环数组
func nextGreaterElements(nums []int) []int {
	res, stk := make([]int, 0), make([]int, 0)
	for i := 0; i < len(nums); i++ {
		res = append(res, -1)
	}

	for i := 0; i < len(nums)*2; i++ {
		num := nums[i%len(nums)]
		for len(stk) > 0 && nums[stk[len(stk)-1]] < num {
			index := stk[len(stk)-1]
			res[index] = num
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, i%len(nums))
	}
	return res
}
