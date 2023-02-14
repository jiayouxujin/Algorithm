package main

func longestValidParentheses(s string) int {
	dp, stack := make([]int, len(s)+1), make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
			dp[i+1] = 0
		} else {
			if len(stack) > 0 {
				leftIndex := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				len := i + 1 - leftIndex + dp[leftIndex]
				dp[i+1] = len
			} else {
				dp[i+1] = 0
			}
		}
	}
	res := 0
	for i := 0; i < len(dp); i++ {
		res = max(res, dp[i])
	}
	return res
}
