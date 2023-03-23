package main

func maxSlidingWindow(nums []int, k int) []int {
	if nums == nil || len(nums) == 0 || k == 0 {
		return nums
	}

	stack, res := make([]int, 0), make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && stack[0] < i-k+1 {
			stack = stack[1:]
		}

		//保持单调
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)

		if i >= k-1 {
			res = append(res, nums[stack[0]])
		}
	}
	return res
}
