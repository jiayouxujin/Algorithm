package main

func dailyTemperatures(temperatures []int) []int {
	res, stack := make([]int, len(temperatures)), make([]int, 0)
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] <= temperatures[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			res[i] = stack[len(stack)-1] - i
		} else {
			res[i] = 0
		}
		stack = append(stack, i)
	}
	return res
}
