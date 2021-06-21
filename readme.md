# 刷题日记

| 日期       | 题目               | 完成情况 |
| ---------- | ------------------ | -------- |
| 2021-06-10 | 81(1) 33(2)        | 👍        |
| 2021-06-11 | 81(2) 242(1)       | 👍        |
| 2021-06-12 | 242(2) 349(1)      | 👍        |
| 2021-06-13 | 349(2)             | 👍        |
| 2021-06-14 |                    |          |
| 2021-06-15 | 202(1)             | 👍        |
| 2021-06-16 | 33(3) 202(2) 1(1)  | 👍        |
| 2021-06-17 | 81(3)1(2)          | 👍        |
| 2021-06-18 | 242(3)             | 👍        |
| 2021-06-19 | 349(3) 1155(1)     | 👍        |
| 2021-06-20 | 1155(2) 474(1)     | 👍        |
| 2021-06-21 | 202(3)474(2)162(1) | 👍        |
| 2021-06-22 | 1(3)162(2)         |          |
|            |                    |          |
|            |                    |          |
|            |                    |          |
|            |                    |          |
|            |                    |          |
|            |                    |          |

[toc]

## 心得

- 二分查找只能用于有序数组中.
- 判断数据存不存在使用HashMap

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

### [33. 搜索旋转排序数组](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)

>整数数组 `nums` 按升序排列，数组中的值 **互不相同** 。
>
>在传递给函数之前，`nums` 在预先未知的某个下标 `k`（`0 <= k < nums.length`）上进行了 **旋转**，使数组变为 `[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]`（下标 **从 0 开始** 计数）。例如， `[0,1,2,4,5,6,7]` 在下标 `3` 处经旋转后可能变为 `[4,5,6,7,0,1,2]` 。
>
>给你 **旋转后** 的数组 `nums` 和一个整数 `target` ，如果 `nums` 中存在这个目标值 `target` ，则返回它的下标，否则返回 `-1` 。

首先如果不考虑时间复杂度的话，那么直接`for`一下就行，这样时间复杂度是O(n)。但是这个时间复杂度面试官一般不会接受。

所以他们会提出我需要时间复杂度为O(logn)的解法。

在数组里搜索数据，且时间复杂度为O(logn)就会想到使用`二分查找`，但是这里的二分查找不同于直接有序的二分。这里的数据只能说是某个区间里有序。

二分的应用场景就是在有序的情况下去找数据，所以无序是无法解决的。上面我们提到旋转排序数组只能在某个区间里有序。

> 解法来了，判断目标数字是否在该区间内，如果在直接舍弃另一边。如果不在则舍弃有序的这一边

为什么肯定会有个区间是有序的，因为这个数组只旋转一次。

```java
public class Solution{
    public int search(int[] nums,int target){
        if(nums==null||nums.length==0){
            return -1;
        }
        int left=0,right=nums.length-1;
        while(left<right){
            int mid=(left+right)/2;
            if(nums[mid]==target){
                return mid;
            }
            
            //判断有序的区间
            if(nums[mid]<nums[right]){
                //目标数字在该区间内
                if(nums[mid]<target&&nums[right]>=target){
                    left=mid+1;
                }else{
                    //目标数字不在该区间内
					right=mid-1;
                }
            }else{
                if(nums[mid]>target&&nums[left]<=target){
                    right=mid-1;
                }else{
                    left=mid+1;
                }
            }
        }
        return nums[left]==target?left:-1;
    }
}
```

#### 二刷心得=

二刷了一下，竟然没有AC，问题出现在我判断有序区间的if改成`nums[mid]>nums[left]`,为什么这样的条件不行？

因为我们计算mid的是向下取整，也就是mid是靠近left的。所以有可能出现mid==left，那么这样你不能说右边一定是有序的，所以导致出问题。

### 	[81. 搜索旋转排序数组 II](https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/)

>已知存在一个按非降序排列的整数数组 `nums` ，数组中的值不必互不相同。
>
>在传递给函数之前，`nums` 在预先未知的某个下标 `k`（`0 <= k < nums.length`）上进行了 **旋转** ，使数组变为 `[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]`（下标 **从 0 开始** 计数）。例如， `[0,1,2,4,4,4,5,6,6,7]` 在下标 `5` 处经旋转后可能变为 `[4,5,6,6,7,0,1,2,4,4]` 。
>
>给你 **旋转后** 的数组 `nums` 和一个整数 `target` ，请你编写一个函数来判断给定的目标值是否存在于数组中。如果 `nums` 中存在这个目标值 `target` ，则返回 `true` ，否则返回 `false` 。

这道题相对上一道题就是数字可以重复，那么就是一下子没法判断哪里是有序的区间。为了解决这个问题，先把重复的数字去掉。`nums[left]==nums[mid]==nums[right]`时左右指针都向中间移动，这样就可以在新的区间内寻找有序数组。

并且要注意这里判断的有序数组还要包括`=`号

```java
public class Solution{
    public boolean search(int[] nums,int target){
        if(nums==null||nums.length==0){
            return false;
        }
        int left=0,right=nums.length-1;
        while(left<right){
            int mid=(left+right)/2;
            if(nums[mid]==target){
                return true;
            }
            //判断是否重复
            if(nums[left]==nums[mid]&&nums[mid]==nums[right]){
                left++;
                right--;
            }else if(nums[mid]<=nums[right]){
                if(nums[right]>=target&&target>nums[mid]){
                    left=mid+1;
                }else{
                    right=mid-1;
                }
            }else{
                if(nums[left]<=target&&target<nums[mid]){
                    right=mid-1;
                }else{
                    left=mid+1;
                }
            }
        }
        return nums[left]==target;
    }
}
```

### [242. 有效的字母异位词](https://leetcode-cn.com/problems/valid-anagram/)

>给定两个字符串 *s* 和 *t* ，编写一个函数来判断 *t* 是否是 *s* 的字母异位词.

来到简单题热热身，这道题其实还算简单。主要有个核心思想，想要O(1)判断一个数在不在，可以使用HashMap.

```java
public class Solution{
    public boolean isAnagram(String s,String t){
        if(s==null||t==null||(s.length()!=t.length())){
            return false;
        }
        HashMap<Character,Integer> map=new HashMap<>();
        for(char c:s.toCharArray()){
            if(map.containeskey(c)){
                map.put(c,map.get(c)+1);
            }else{
                map.put(c,1);
            }
        }
        
        for(char c:t.toCharArray()){
            if(!map.containsKey(c)||map.get(c)==0){
                return false;
            }else{
                map.put(c,map.get(c)-1);
            }
        }
        
        return true;
    }
}
```

### [349. 两个数组的交集](https://leetcode-cn.com/problems/intersection-of-two-arrays/)

>给定两个数组，编写一个函数来计算它们的交集。

这也是个简单题，思路就是使用set，这里不使用hashMap的原因是数组的个数很多，可能会导致hashMap很大.

```java
class Solution{
    public int[] intersection(int[] nums1,int[] nums2){
        if(nums1==null||nums1.length==0||nums2==null||nums2.length==0){
            return new int[0];
        }
        
        Set<Integer> set=new HashSet<>();
        Set<Integer> resSet=new HashSet<>();
        
        for(int i:nums1){
            set.add(i);
        }
        for(int i:nums2){
            if(set.contains(i)){
                resSet.add(i);
            }
        }
        
        int[] res=new int[resSet.size()];
        int index=0;
        for(int i:resSet){
            res[index++]=i;
        }
        return res;
    }
}
```

### [202. 快乐数](https://leetcode-cn.com/problems/happy-number/)

>编写一个算法来判断一个数 `n` 是不是快乐数。
>
>「快乐数」定义为：
>
>- 对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
>- 然后重复这个过程直到这个数变为 1，也可能是 **无限循环** 但始终变不到 1。
>- 如果 **可以变为** 1，那么这个数就是快乐数。
>
>如果 `n` 是快乐数就返回 `true` ；不是，则返回 `false` 。

注意题目说如果不会变成1，则结果会是无限循环。所以可以使用set来存储。

```java
class Solution{
    public boolean isHappy(int n){
        HashSet<Integer> set=new HashSet<>();
        while(n!=1){
            if(set.contains(n)){
                return false;
            }
            set.add(n);
            n=getNextNum(n);
        }
        return n==1;
    }
    
    public int getNextNum(int n){
        int sum=0;
        while(n>0){
            int tmp=n%10;
            sum+=tmp*tmp;
            n=n/10;
        }
        return sum;
    }
}
```

### [1. 两数之和](https://leetcode-cn.com/problems/two-sum/)

>给定一个整数数组 `nums` 和一个整数目标值 `target`，请你在该数组中找出 **和为目标值** *`target`* 的那 **两个** 整数，并返回它们的数组下标。
>
>你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
>
>你可以按任意顺序返回答案。

这道题就是可以使用HashMap来减少遍历次数。只需要遍历一次就可以。从而使得时间复杂度是O(n)

```java
class Solution{
    public int[] twoSum(int[] nums,int target){
        HashMap<Integer,Integer> map=new HashMap<>();
        for(int i=0;i<nums.length;i++){
            int tmp=target-nums[i];
            if(map.containsKey(tmp)){
                int[] res=new int[2];
                res[0]=map.get(tmp);
                res[1]=i;
                return res;
            }
            map.put(nums[i],i);
        }
        return new int[0];
    }
}
```

### [1155. 掷骰子的N种方法](https://leetcode-cn.com/problems/number-of-dice-rolls-with-target-sum/)

>这里有 d 个一样的骰子，每个骰子上都有 f 个面，分别标号为 1, 2, ..., f。
>
>我们约定：掷骰子的得到总点数为各骰子面朝上的数字的总和。
>
>如果需要掷出的总点数为 target，请你计算出有多少种不同的组合情况（所有的组合情况总共有 f^d 种），模 10^9 + 7 后返回。

动态规划题。

先找到题目中变化的点，即骰子的个数和当前的总和。

我们设为dp-i-j，其中i表示第i个骰子，j表示当前的总和。现在我们可以写出动态规划的公式

因为第i个骰子只能扔出1-f的数，所以dp-i-j=dp【i-1】【j-1】+....dp【i-1】【j-f】

接下来base，就是只有一个骰子扔出1-f的组合情况只有1种.[因为只能扔一次]

```java
class Solution{
    public int numRollsToTarget(int d,int f,int target){
        int mod=1000000007;
        int[][] dp=new int[31][1001];
        int min=Math.min(f,target);
        for(int i=1;i<=min;i++){
            dp[1][i]=1;
        }
        int maxTarget=d*f;
        for(int i=2;i<=d;i++){
            for(int j=i;j<=maxTarget;j++){
                for(int k=1;j-k>0&&k<=f;k++){
                    dp[i][j]=(dp[i][j]+dp[i-1][j-k])%mod;
                }
            }
        }
        return dp[d][target];
    }
}
```

### [474. 一和零](https://leetcode-cn.com/problems/ones-and-zeroes/)

> 给你一个二进制字符串数组 strs 和两个整数 m 和 n 。 请你找出并返回 strs 的最大子集的大小，该子集中 最多 有 m 个 0 和 n 个 1 。 如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集

这是一个动态规划题，可以变化的东西是0的个数和1的个数，所以我们设DP[i][j]表示当前是i个0,j个1有多少个元素。

这道题需要注意的一个地方是它是从一个二进制字符串数组里寻找。所以我们开始遍历这个strs，然后dp[i][j]=max(dp[i][j],dp[i-zero][j-one]+1)

解释一下这个DP公式，zero,one分别表示当前str的0和1的个数。

```java
class Solution {
    public int findMaxForm(String[] strs, int m, int n) {
    int[][] dp=new int[m+1][n+1];	
    //base
    dp[0][0]=0;
    for(String str:strs){
		int[] zeroAndOne=count(str);
		int zero=zeroAndOne[0];
		int one=zeroAndOne[1];
        for(int i=m;i>=zero;i--){
			for(int j=n;j>=one;j--){
		        dp[i][j]=Math.max(dp[i][j],dp[i-zero][j-one]+1);		
            }
        }
    }
    return dp[m][n];
    }  
   public int[] count(String str){
		int[] res=new int[2];
        for(char c:str.toCharArray()){
			if(c=='0')
				res[0]++;
			else
				res[1]++;
	    }
    return res;
    }
}
```

