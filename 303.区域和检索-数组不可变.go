/*
 * @lc app=leetcode.cn id=303 lang=golang
 *
 * [303] 区域和检索 - 数组不可变
 */

// @lc code=start
type NumArray struct {
	prefixSum []int
}


func Constructor(nums []int) NumArray {
	for i:=1;i<len(nums);i++{
		nums[i]+=nums[i-1]
	}
	return NumArray{prefixSum:nums}
}


func (this *NumArray) SumRange(left int, right int) int {
	if left>0{
		return this.prefixSum[right]-this.prefixSum[left-1]
	}
	return this.prefixSum[right]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
// @lc code=end

