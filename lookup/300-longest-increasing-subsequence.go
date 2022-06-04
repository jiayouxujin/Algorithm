package main

func lengthOfLIS(nums []int) int {
	dp, res := make([]int, len(nums)), 0
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	for j := 1; j < len(nums); j++ {
		for i := j - 1; i >= 0; i-- {
			if nums[j] > nums[i] {
				dp[j] = max(dp[j], dp[i]+1)
			}
		}
	}
	for i := 0; i < len(dp); i++ {
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

//func main() {
//	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
//	fmt.Printf("%v \n", lengthOfLIS(nums))
//}
