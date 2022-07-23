package main

func minDistance2(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			down := dp[i-1][j] + 1
			left := dp[i][j-1] + 1
			leftDown := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				leftDown += 1
			}

			dp[i][j] = min(down, min(leftDown, left))
		}
	}
	return dp[m][n]
}
