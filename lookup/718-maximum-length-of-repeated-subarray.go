package main

func findLength(nums1 []int, nums2 []int) int {
	res, dp := 0, make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}
	for i := len(nums1) - 1; i >= 0; i-- {
		for j := len(nums2) - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}
	return res
}

//func main() {
//	nums1, nums2 := []int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}
//	fmt.Printf("%v \n", findLength(nums1, nums2))
//}
