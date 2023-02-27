package main

//func numWays(n int) int {
//	if n <= 1 {
//		return 1
//	}
//	dp, mod := make([]int, n+1), 1000000007
//	dp[0], dp[1] = 1, 1
//	for i := 2; i <= n; i++ {
//		dp[i] = (dp[i-1] + dp[i-2]) % mod
//	}
//	return dp[n]
//}

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	a, b := 1, 1
	for i := 2; i <= n; i++ {
		a, b = b, (a+b)%1000000007
	}
	return b
}
