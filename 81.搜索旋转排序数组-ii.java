/*
 * @lc app=leetcode.cn id=81 lang=java
 *
 * [81] 搜索旋转排序数组 II
 */

// @lc code=start
class Solution81 {
    public boolean search(int[] nums, int target) {
        if (nums == null || nums.length == 0) {
            return false;
        }

        int left = 0, right = nums.length - 1;
        while (left < right) {
            int mid = (left + right) / 2;
            if (nums[mid] == target) {
                return true;
            }
            if (nums[left] == nums[mid] && nums[mid] == nums[right]) {
                left++;
                right--;
            } else {
                if (nums[mid] <= nums[right]) {
                    if (target > nums[mid] && target <= nums[right]) {
                        left = mid + 1;
                    } else {
                        right = mid - 1;
                    }
                } else {
                    if (target >= nums[left] && target < nums[mid]) {
                        right = mid - 1;
                    } else {
                        left = mid + 1;
                    }
                }
            }
        }
        return nums[left] == target;
    }

    public static void main(String args[]) {
        Solution81 solution = new Solution81();
        int[] nums = new int[] { 1, 1, 1, 1, 1, 1, 1, 1, 1, 13, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1 };
        solution.search(nums, 13);
    }
}
// @lc code=end
