package main

func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}

	sz := len(prices)
	dp := make([][]int, sz)
	for i := 0; i < sz; i++ {
		dp[i] = make([]int, 2)
	}
	// 0 表示卖出去,1表示买入
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][1]+prices[i], dp[i-1][0])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	return dp[sz-1][0]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
