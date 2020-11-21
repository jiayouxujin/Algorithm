/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 */

// @lc code=start
func maxProduct(nums []int) int {
	minimum,maximum,res:=nums[0],nums[0],nums[0]
	for i:=1;i<len(nums);i++{
		if nums[i]<0{
			maximum,minimum=minimum,maximum
		}
		maximum=max(nums[i],maximum*nums[i])
		minimum=min(nums[i],minimum*nums[i])
		res=max(res,maximum)
	}
	return res
}
func max(a,b int) int {
	if a>=b{
		return a
	}
	return b
}

func min(a,b int) int {
	if a<=b{
		return a
	}
	return b
}
// @lc code=end

