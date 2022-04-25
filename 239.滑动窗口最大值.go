/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums)==0||len(nums)<k{
		return make([]int, 0)
	}

	window:=make([]int,0,k)
	result:=make([]int,0,len(nums)-k+1)

	for i,v:=range nums{
		if len(window)>0&&window[0]<=i-k{
			window=window[1:len(window)]
		}

		for len(window)>0&&nums[window[len(window)-1]]<v{
			window=window[0:len(window)-1]
		}

		window=append(window,i)
		if i>=k-1{
			result=append(result,nums[window[0]])
		}
	}
	return result
}
// @lc code=end

