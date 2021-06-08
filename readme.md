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

------
#### 分割线，最近由于总总原因断了刷题，要重新拾起来。接下来的刷题策略是按照专题来进行

### Array
| 题目  | 时间复杂度     | 空间复杂度  | 完成时间 |
| ----- | ----------- | -------- | -------- |
| 1.两数之和 | O(n) | O(n) | 2020-12-12 |
| 11.盛最多水的容器 | O(n) | O(1) | 2020-12-13 |

### Tree
| 题目  | 时间复杂度     | 空间复杂度  | 完成时间 |
| ----- | ----------- | -------- | -------- |
| 94.二叉树的中序遍历 | O(n) | O(1) | 2020-12-20 |
| 111.二叉树的最小深度 | O(n) | O(1) | 2020-12-20 |
| 112.路径总和 | O(n) | O(1) | 2020-12-19 |
| 113.路径总和2 | O(n) | O(1) | 2020-12-20 |
| 129.求根到叶子节点数字之和 | O(n) | O(1) | 2020-12-20 |



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
### leetcode 1
> 经典的两数之和，用上Hash这个思想和巧妙
### leetcode 11
> 经典双指针的思想-盛最多水的容器

通过left和right。如果left小于等于right则移动left。这里证明为什么这样可以：

如果不可以的话，意味着有可能还会用到left(这里假设left为小的)

那么此时不管right是大的还是小的都不会比原来得到的结果要大

```java
public class Solution{
    public int solve(int []height){
        int left=0,right=height.length-1;
        int ans=0;
        while(left<=right){
            ans=Math.max(ans,Math.min(height[left],height[right])*(right-left));
            if(height[left]<=height[right]){
                left++;
            }else{
                right--;
            }
        }
        return ans;
    }
}
```

### Leetcode 42

>给定 *n* 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

单调栈，为什么会是单调栈的解法呢？木桶原理。所以我们维护一个单调的序列，单出现不单调的那么就存在可以接水的地方。

```java
public class Solution{
    public int solve(int []height){
        Stack<Integer> stack=new Stack<>();
       	int ans=0;
        
        for(int i=0;i<height.length;i++){
            while(!stack.isEmpty()&&(height[statck.peek()]<=height[i])){
                int popIndex=stack.pop();
                if(!stack.isEmpty()){
                int h=Math.min(height[i]-height[popIndex],height[stack.peek()]-height[popIndeX])；
                int w=i-stack.peek()-1;
                    
                ans+=h*w;
            }
            }
            stack.push(i);
        }
        return ans;
    }
}
```

### Leetcode 24 两两交换链表中的节点

>给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
>
>**你不能只是单纯的改变节点内部的值**，而是需要实际的进行节点交换。

这应该是一道模拟题。有一种方法是通过递归。因为原问题是两两交换，子问题也是两两交换。

```java
publc class Solution{
    public ListNode swapPairs(ListNode head){
        //base
        if(head==null||head.next==null) return head;
        
        //process
        ListNode next=head.next;
        //cur next
        ListNode newNode=swapPairs(next.next);
        
        //process
        next.next=head;
        head.next=newNode;
        
        return next;
    }
}
```

第二个方法就是模拟，直接来。在做的过程中有个误区就是next不知道是啥。后面发现如果是node.next.next 表示的是node.next所指定的node的next属性，不是node

```java
public class Solution{
    public ListNode swapPairs(ListNode head){
        ListNode dummy=new ListNode();
        dummy.next=head;
        ListNode pre=dummpyHead;
        
        while(pre.next!=null&&pre.next.next!=null){
            ListNode tmp=head.next.next;
            pre.next=head.next;
            head.next.next=head;
            head.next=tmp;
            
            pre=head;
            head=tmp;
        }
        return dummy.next;
    }
}
```

### Leetcode 19 删除链表的倒数第N个节点

>给你一个链表，删除链表的倒数第 `n` 个结点，并且返回链表的头结点。

快慢指针，不过为了解决头指针被删掉的情况需要加一个虚拟头指针。

```java
public class Solution{
    public ListNode removeNthFromEnd(ListNode head,int n){
        ListNode dummy=new ListNode();
        dummy.next=head;
        ListNode slow=dummy,fast=dummy;
        while(n>=0){
            fast=fast.next;
            n--;
        }
        
        while(fast!=null){
            fast=fast.next;
            slow=slow.next;
        }
        
        slow.next=slow.next.next;
        return dummy.next;
    }
}
```

### [面试题 02.07. 链表相交](https://leetcode-cn.com/problems/intersection-of-two-linked-lists-lcci/)

>给定两个（单向）链表，判定它们是否相交并返回交点。请注意相交的定义基于节点的引用，而不是基于节点的值。换句话说，如果一个链表的第k个节点与另一个链表的第j个节点是同一节点（引用完全相同），则这两个链表相交。

这道题题目一开始没看懂，不过一开始的思路是用一个container存起来，然后遍历另外一个链表，如果存在则说明相交。但是这样的空间复杂度就变成了O(n)。最终参考答案，才发现题目是说如果两个链表相交就是说最终会合成同一条。所以就有这三种情况

- 尾部不同不相交
- 一个为空不相交
- 相交

```java
public class Solution{
    public ListNode getIntersectionNode(ListNode headA,ListNode headB){
        if(headA==null||headB==null) return null;
        ListNode a=headA,b=headB;
        int sizea=0,sizeb=0;
        while(a!=null){
            a=a.next;
            sizea++;
        }
        while(b!=null){
            b=b.next;
            sizeb++;
        }
        int diff=Math.abs(sizea-sizeb);
        if(sizea>sizeb){
            while(diff-->0){
                headA=headA.next;
            }
        }else{
            while(diff-->=){
                headB=headB.next;
            }
        }
        while(headA!=headB){
            headA=headA.next;
            headB=headB.next;
        }
        return headA;
    }
}
```

### [142. 环形链表 II](https://leetcode-cn.com/problems/linked-list-cycle-ii/)

>给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
>
>为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。

首先要通过快慢指针判断是否存在环。然后通过公式发现找到入口的窍门

x(到入口) y(慢指针在环内走的)z(环的长度-y)

2*(x+y)=x+y+n(y+z)  //x+y是慢指针走的，x+n(y+z)是快指针走的。慢指针比快指针少走一半

x+y=n(y+z)

x=(n-1)(y+z)-y

x=z

```java
public class Solution{
    public ListNode detectCycle(ListNode head){
        ListNode fast=head,slow=head;
        while(fast!=null&&fast.next!=null){
            slow=slow.next;
            fast=fast.next.next;
            
            if(slow==fast){
                ListNode index1=head;
                ListNode index2=slow;
                while(index1!=index2){
                    index1=index1.next;
                    index2=index2.next;
                }
                return index1;
            }
        }
        return null;
    }
}
```



