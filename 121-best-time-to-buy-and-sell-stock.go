package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	sz := len(prices)
	dp := make([][]int, sz)
	for i := 0; i < sz; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < sz; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	return dp[sz-1][0]
}

//func main() {
//	prices := []int{7, 1, 5, 3, 6, 4}
//	fmt.Printf("%v \n", maxProfit(prices))
//}
