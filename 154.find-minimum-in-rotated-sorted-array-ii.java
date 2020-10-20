/*
 * @lc app=leetcode id=154 lang=java
 *
 * [154] Find Minimum in Rotated Sorted Array II
 */

// @lc code=start
class Solution {
    public int findMin(int[] nums) {
       int left=0,right=nums.length-1;
       while(left<right){
           int mid=left+(right-left)/2;
           if(nums[mid]>nums[right]) left=mid+1;
           else if(nums[mid]<nums[right]) right=mid;
           else right--;
       }
       return nums[left];
    }
}
// @lc code=end

