package main

func longestOnes(nums []int, k int) int {
	left, right, res, modify := 0, 0, 0, 0
	for right < len(nums) {
		n := nums[right]
		right++
		if n == 0 {
			modify++
		}
		if modify > k {
			if nums[left] == 0 {
				modify--
			}
			left++
		}
		res = max(res, right-left)
	}
	return res
}
