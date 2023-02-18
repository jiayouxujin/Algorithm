package main

func lengthOfLIS(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 1

	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
	}
	res := 0
	for i := 0; i < len(dp); i++ {
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
