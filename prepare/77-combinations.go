package main

func combine(n int, k int) [][]int {
	var res [][]int
	backtraceCombine(n, 1, k, []int{}, &res)
	return res
}

func backtraceCombine(n, start, k int, cur []int, res *[][]int) {
	if len(cur) == k {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}
	for i := start; i <= n; i++ {
		cur = append(cur, i)
		backtraceCombine(n, i+1, k, cur, res)
		cur = cur[:len(cur)-1]
	}
}
