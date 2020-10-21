/*
 * @lc app=leetcode id=74 lang=java
 *
 * [74] Search a 2D Matrix
 */

// @lc code=start
class Solution {
    public boolean searchMatrix(int[][] matrix, int target) {
        if(matrix==null||matrix.length==0){
            return false;
        }
        int n=matrix[0].length;
        int left=0,right=n*matrix.length-1;
        while(left<=right){
            int mid=left+(right-left)/2;
            if(matrix[mid/n][mid%n]==target){
                return true;
            }else if(matrix[mid/n][mid%n]>target){
                right=mid-1;
            }else{
                left=mid+1;
            }
        }
        return false;
    }
}
// @lc code=end

