/**
 * Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.
 *
 * Your algorithm's runtime complexity must be in the order of O(log n).
 *
 * If the target is not found in the array, return [-1, -1].
 */
class Solution {
    public int[] searchRange(int[] nums, int target) {
        int num[]=new int[2];
        int ans=-1;
        if(nums.length!=0){
            ans=search(nums,0,nums.length-1,target);
        }
        int i=-1,j=-1;
        if(ans!=-1){
            i=ans;
            j=ans;
            while(i>=0&&nums[i]==target) i--;
            while(j<nums.length&&nums[j]==target) j++;
            i+=1;
            j-=1;
        }
        num[0]=i;
        num[1]=j;
        return num;
    }

    private int search(int[] nums, int begin, int end, int target) {
        //判断不可能找到
        if(begin>end||nums[begin]>target||nums[end]<target){
            return -1;
        }
        //二分查找
        int mid;
        while(begin<=end){
            mid=begin+((end-begin)>>1);
            if(nums[mid]==target) return mid;
            else if(nums[mid]>target){
                end=mid-1;
            }else{
                begin=mid+1;
            }
        }
        return -1;
    }

}