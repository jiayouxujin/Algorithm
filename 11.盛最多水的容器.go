package algorithm
/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	max,start,end:=0,0,len(height)-1

	for start<end{
		width:=end-start
		h:=0
		
		if height[start]<height[end]{
			h=height[start]
			start++
		}else{
			h=height[end]
			end--
		}

		tmp:=h*width
		if tmp>max{
			max=tmp
		}
	}
	return max
}
// @lc code=end

