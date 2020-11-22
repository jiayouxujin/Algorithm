# 刷题日记

## 定个小目标，每日一题

| 日期  | 题目        | 解决情况 |
| ----- | ----------- | -------- |
| 10/18 | leetcode 81 | done     |
| 10/19 | leetcode 153 | done     |
| 10/20 | leetcode 154 | done     |
| 10/21 | leetcode 74 | done     |
| 10/23 | leetcode 547 | done     |
| 10/24 | leetcode 721 | done     |
| 10/25 | leetcode 104 | done     |
| 10/26 | leetcode 112 | done     |
| 10/27 | leetcode 113 | done     |
| 10/28 | leetcode 130 | done     |
| 10/29 | leetcode 200 | done     |
| 10/30 | leetcode 695 | done     |
| 10/31 | leetcode 979 | done     |
| 11/01 | leetcode 22 | done     |
| 11/02 | leetcode 51 |      |
| 11/03 | leetcode 52 |    |
| 11/04 | leetcode 37 |      |
| 11/05 | leetcode 980 | done     |
| 11/06 | leetcode 79 | done     |
| 11/07 | leetcode 240 | done     |
| 11/08 | leetcode 23 | done     |
| 11/09 | leetcode 257 | done     |
| 11/10| leetcode  |      |
| 11/11| leetcode 973 |   done   |
| 11/12| leetcode  |      |
| 11/13| leetcode  |      |
| 11/14| leetcode 70 |   done   |
| 11/21| leetcode 152 |   done   |
| 11/22| leetcode 53 |   done   |



## 题解
### leetcode23

>Merge *k* sorted linked lists and return it as one sorted list. Analyze and describe its complexity.
>
>**Example:**
>
>```
>Input:
>[
>  1->4->5,
>  1->3->4,
>  2->6
>]
>Output: 1->1->2->3->4->4->5->6
>```

解题思路：二分，两两进行合并。

因为这道题有重复性，100合并跟2组50个合并是一致的所以可以使用二分。

【如何提高性能？用线程---来自某次面试】

### leetcode257

>Given a binary tree, return all root-to-leaf paths.
>
>**Note:** A leaf is a node with no children.
>
>**Example:**
>
>```
>Input:
>
>   1
> /   \
>2     3
> \
>  5
>
>Output: ["1->2->5", "1->3"]
>
>Explanation: All root-to-leaf paths are: 1->2->5, 1->3
>```

解题思路：递归，因为是树的问题，所以容易想到递归。并且这里存在重复问题，可以使用递归。

### leetcode52

>Given an integer array `nums`, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
>
>**Example:**
>
>```
>Input: [-2,1,-3,4,-1,2,1,-5,4],
>Output: 6
>Explanation: [4,-1,2,1] has the largest sum = 6.
>```

解题思路：DP

dp[i]表示前i个最大的和，那么存在两种情况

dp[i-1]<0 那么dp[i]=nums[i]

dp[i-1]>0 那么dp[i]=dp[i-1]+nums[i]