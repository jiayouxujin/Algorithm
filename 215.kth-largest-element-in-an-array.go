/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
    if len(nums)==0{
		return 0
	}
	return selection(nums,0,len(nums)-1,len(nums)-k)
}

func selection(nums []int,l,r,k int)int{
	if l==r{
		return nums[l]
	}
	
	p:=r
	ll:=l
	rr:=r-1
	for ll<=rr{
		for(ll<=rr&&nums[ll]<=nums[p]){
			ll++
		}
		for(ll<=rr&&nums[rr]>nums[p]){
			rr--
		}

		if(ll<=rr){
			nums[ll],nums[rr]=nums[rr],nums[ll]
			ll++
			rr--
		}
	}
	nums[ll],nums[p]=nums[p],nums[ll]
	p=ll
	if(p>k-1){
		return selection(nums,l,p-1,k)
	}else if(p<k-1){
		return selection(nums,p+1,r,k-p)
	}
	return nums[p-1];
}
// @lc code=end

