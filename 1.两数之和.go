/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	var m=make(map[int]int)

	for index,num:=range nums{
		if value,ok:=m[target-num];ok{
			return []int{index,value}
		}
		m[num]=index
	}
	return nil
}
// @lc code=end

