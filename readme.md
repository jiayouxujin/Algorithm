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