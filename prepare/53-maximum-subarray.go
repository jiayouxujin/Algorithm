package main

func maxSubArray(nums []int) int {
	dp, res := make([]int, len(nums)), 0
	dp[0], res = nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = nums[i]
		if dp[i-1]+nums[i] > dp[i] {
			dp[i] = dp[i-1] + nums[i]
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
