package main

func generateParenthesis(n int) []string {
	var res []string
	backtraceParenthesis("", 0, 0, n, &res)
	return res
}

func backtraceParenthesis(cur string, left, right, n int, res *[]string) {
	if left == n && right == n {
		*res = append(*res, cur)
		return
	}
	if left < n {
		cur += "("
		backtraceParenthesis(cur, left+1, right, n, res)
		cur = cur[:len(cur)-1]
	}

	if right < left {
		cur += ")"
		backtraceParenthesis(cur, left, right+1, n, res)
		cur = cur[:len(cur)-1]
	}
}
