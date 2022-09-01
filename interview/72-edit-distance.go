package main

func minDistance(word1 string, word2 string) int {
	len1, len2 := len(word1), len(word2)
	dp := make([][]int, len1+1)
	for i := 0; i <= len1; i++ {
		dp[i] = make([]int, len2+1)
	}

	for i := 1; i <= len1; i++ {
		dp[i][0] = i
	}

	for j := 1; j <= len2; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			left := dp[i][j-1] + 1
			down := dp[i-1][j] + 1
			leftdown := dp[i-1][j-1]

			if word1[i-1] != word2[j-1] {
				leftdown += 1
			}

			dp[i][j] = min(left, min(down, leftdown))
		}
	}
	return dp[len1][len2]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
