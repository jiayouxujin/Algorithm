package main

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m*n == 0 {
		return m + n
	}
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	//init
	for i := 0; i < m+1; i++ {
		dp[i][0] = i
	}
	for i := 0; i < n+1; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			left := dp[i-1][j] + 1
			down := dp[i][j-1] + 1
			left_down := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				left_down += 1
			}
			dp[i][j] = min(left_down, min(left, down))
		}
	}
	return dp[m][n]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

//func main() {
//	word1, word2 := "horse", "ros"
//	fmt.Printf("%v \n", minDistance(word1, word2))
//}
