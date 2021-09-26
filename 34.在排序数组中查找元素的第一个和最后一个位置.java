/*
 * @lc app=leetcode.cn id=34 lang=java
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
class Solution {
    public int[] searchRange(int[] nums, int target) {
        if (nums == null || nums.length == 0) {
            return new int[] { -1, -1 };
        }
        int left = 0, right = nums.length - 1;
        while (left <= right) {
            int mid = (left + right) / 2;
            if (nums[mid] == target) {
                int tmpLeft = mid, tmpRight = mid;
                while (tmpLeft >= 0 && nums[tmpLeft] == target) {
                    tmpLeft--;
                }
                while (tmpRight < nums.length && nums[tmpRight] == target) {
                    tmpRight++;
                }
                return new int[] { tmpLeft + 1, tmpRight - 1 };
            }
            if (nums[mid] < target) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        return new int[]{-1,-1};
    }
}
// @lc code=end
