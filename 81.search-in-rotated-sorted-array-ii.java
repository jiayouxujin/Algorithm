/*
 * @lc app=leetcode id=81 lang=java
 *
 * [81] Search in Rotated Sorted Array II
 */

// @lc code=start
class Solution {
    public boolean search(int[] nums, int target) {
        int left=0,right=nums.length-1;

        while(left<=right){
            int mid=left+(right-left)/2;
            if(target==nums[mid]){
                return true;
            }else if(nums[left]==nums[mid]){
                left++;
            }else if(nums[left]<nums[mid]){
                if(nums[left]<=target&&target<nums[mid]){
                    right=mid-1;
                }else{
                    left=mid+1;
                }
            }else{
                if(nums[mid]<target&&target<=nums[right]){
                    left=mid+1;
                }else{
                    right=mid-1;
                }
            }
        }
        return false;
    }
}
// @lc code=end

