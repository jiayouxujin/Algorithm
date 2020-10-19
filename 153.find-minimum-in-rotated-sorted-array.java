/*
 * @lc app=leetcode id=153 lang=java
 *
 * [153] Find Minimum in Rotated Sorted Array
 */

// @lc code=start
class Solution {
    public int findMin(int[] nums) {
        if(nums==null||nums.length==0){
            return 0;
        }
        if(nums.length==1){
            return nums[0];
        }


        int left=0,right=nums.length-1;
        while(left<right){
            int mid=left+(right-left)/2;
            if(mid>0&&nums[mid]<nums[mid-1]){
                return nums[mid];
            }
            if(nums[left]<=nums[mid]&&nums[mid]>nums[right]){
                left=mid+1;
            }else{
                right=mid-1;
            }
        }
        return nums[left];
    }
}
// @lc code=end

