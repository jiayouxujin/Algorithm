package main

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = maxMaxSubArray(nums[i], dp[i-1]+nums[i])
	}

	res := dp[0]
	for i := 0; i < len(dp); i++ {
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

func maxMaxSubArray(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

//func main() {
//	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
//	fmt.Printf("%v \n", maxSubArray(nums))
//}
