package main

func longestValidParentheses(s string) int {
	stk, dp := make([]int, 0), make([]int, len(s)+1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stk = append(stk, i)
			dp[i+1] = 0
		} else {
			if len(stk) > 0 {
				leftIndex := stk[len(stk)-1]
				stk = stk[:len(stk)-1]

				len := 1 + i - leftIndex + dp[leftIndex]
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
