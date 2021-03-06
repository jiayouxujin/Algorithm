# 刷题日记

| 日期       | 题目               | 完成情况 |
| ---------- | ------------------ | -------- |
| 2021-07-01 | 433(1)             | 👍        |
| 2021-07-02 | 1020(1)433(2)      | 👍        |
| 2021-07-03 | 1020(2)            | 👍        |
| 2021-07-04 | 1091(1)            | 👍        |
| 2021-07-05 | 165(1)             | 👍        |
| 2021-07-06 | 剑指offer 57(1)    | 👍        |
| 2021-07-08 | 1367(1)            | 👍        |
| 2021-07-09 | 821(1)             | 👍        |
| 2021-07-10 | 881(1)             | 👍        |
| 2021-07-11 | 532(1)             | 👍        |
| 2021-07-12 | 599(1)             | 👍        |
| 2021-07-13 | 面试题16.24 数对和 | 👍        |
| 2021-07-14 | block              |          |
| 2021-07-15 | 1742               | 👍        |
| 2021-07-16 | 1487               | 👍        |
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

### [162. 寻找峰值](https://leetcode-cn.com/problems/find-peak-element/)

>峰值元素是指其值大于左右相邻值的元素。
>
>给你一个输入数组 `nums`，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 **任何一个峰值** 所在位置即可。
>
>你可以假设 `nums[-1] = nums[n] = -∞` 。

这道题有个限制就是nums[i]≠nums[i+1]。

接着我们来讨论一个峰值可能出现的位置。如果数组是呈现递增的趋势，那么峰值在右边。如果数组呈现递减的趋势那么峰值在左边。

然后呢，因为我们只需要返回一个峰值即可。二分查找要如何应用到该场景呢？二分查找我们一般是在有序数组里查找某个数。在这道题我们可以将mid跟mid+1进行比较

- 如果nums[mid]>nums[mid+1]说明是呈现递减的趋势，那么峰值在左边.
- 如果nums[mid]<nums[mid+1]说是呈现递增的趋势，那么峰值在右边

```java
class Solution{
    public int findPeakElement(int[] nums){
        return search(nums,0,nums.length-1);
    }
    
    public int search(int[] nums,int left,int right){
        if(left==right){
            return left;
        }
        int mid=(left+right)/2;
        if(nums[mid]>nums[mid+1]){
            return search(nums,left,mid);
        }
        return search(nums,mid+1,right);
    }
}
```

### [658. 找到 K 个最接近的元素](https://leetcode-cn.com/problems/find-k-closest-elements/)

>给定一个排序好的数组 `arr` ，两个整数 `k` 和 `x` ，从数组中找到最靠近 `x`（两数之差最小）的 `k` 个数。返回的结果必须要是按升序排好的。
>
>整数 `a` 比整数 `b` 更接近 `x` 需要满足：
>
>- `|a - x| < |b - x|` 或者
>- `|a - x| == |b - x|` 且 `a < b`

这道题比较容易想到的方法就是O(n)

方法1：使用Collection.sort，相当于重新编写排序的规则

方法2：双指针，left和right往中间移动，绝对值大的一方要移动

```java
class Solution{
    public List<Integer> findClosestElements(int[] arr,int k,int x){
        int remove=arr.length-k;
        int left=0,right=arr.length-1;
        while(remove>0){
            if(Math.abs(arr[left]-x)>Math.abs(arr[right]-x))
                left++;
            else
                right--;
            remove--;
        }
        List<Integer> res=new ArrayList<>();
        for(int i=left;i<left+k;i++){
            res.add(arr[i]);
        }
        return res;
    }
}
```

### [744. 寻找比目标字母大的最小字母](https://leetcode-cn.com/problems/find-smallest-letter-greater-than-target/)

>给你一个排序后的字符列表 `letters` ，列表中只包含小写英文字母。另给出一个目标字母 `target`，请你寻找在这一有序列表里比目标字母大的最小字母。
>
>在比较时，字母是依序循环出现的。举个例子：
>
>- 如果目标字母 `target = 'z'` 并且字符列表为 `letters = ['a', 'b']`，则答案返回 `'a'`

这道题有两个思路。一个是直接遍历，还有一个是二分查找，不过这里的二分查找不是寻找相等的，寻找最接近的那个。

```java
class Solution{
    public char nextGreatestLetter(char[] letters,char target){
        int left=0,right=letters.length-1;
        if(letters[right]<=target||letters[left]>target)
            return letters[left];
        while(left+1<right){
            int mid=(left+right)/2;
            if(letters[mid]>target)
                right=mid;
            else
                left=mid;
        }
        return letters[right];
    }
}
```

### [875. 爱吃香蕉的珂珂](https://leetcode-cn.com/problems/koko-eating-bananas/)

>珂珂喜欢吃香蕉。这里有 `N` 堆香蕉，第 `i` 堆中有 `piles[i]` 根香蕉。警卫已经离开了，将在 `H` 小时后回来。
>
>珂珂可以决定她吃香蕉的速度 `K` （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 `K` 根。如果这堆香蕉少于 `K` 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。 
>
>珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。
>
>返回她可以在 `H` 小时内吃掉所有香蕉的最小速度 `K`（`K` 为整数）。

这道题竟然是一个二分法。二分查找的数是K，区间是[1,maxval]。其中maxval是piles中的最大值。然后每次判断cal(piles,k)与H的关系。

```java
class Solution{
    public int minEatingSpeed(int[] piles,int h){
        int maxVal=1;
        for(int pile:piles){
            maxVal=Math.max(maxVal,pile);
        }
        int left=1,right=maxVal;
        while(left<right){
            int mid=(left+right)/2;
            if(cal(piles,mid)>h){
                left=mid+1;
            }else{
                right=mid;
            }
        }
        return right;
    }
    public int cal(int[] piles,int k){
        int sum=0;
        for(int pile:piles){
            sum+=(pile+k-1)/k;
        }
        return sum;
    }
}
```

### [441. 排列硬币](https://leetcode-cn.com/problems/arranging-coins/)

>你总共有 n 枚硬币，你需要将它们摆成一个阶梯形状，第 k 行就必须正好有 k 枚硬币。
>
>给定一个数字 n，找出可形成完整阶梯行的总行数。
>
>n 是一个非负整数，并且在32位有符号整型的范围内。

这也是一道二分法的题目。left=0,right=n。然后通过n(n+1)/2来计算等差数列的总共。

```java
public Solution{
    public int arrangeCoins(int n){
        long left=0,right=n;
        while(left<right){
            long mid=left+(right-left+1)/2;
            long sum=(1+mid)*mid/2;
            if(sum>n){
                right=mid-1;
            }else{
                left=mid;
            }
        }
        return (int)left;
    }
}
```

### [719. 找出第 k 小的距离对](https://leetcode-cn.com/problems/find-k-th-smallest-pair-distance/)

>给定一个整数数组，返回所有数对之间的第 k 个最小**距离**。一对 (A, B) 的距离被定义为 A 和 B 之间的绝对差值。

这也是一道二分法，距离的范围是0-nummax-nummin。然后开始二分。如果count大于k表示此时的mid太大。否则mid太小.

```java
class Solution{
    public int smallestDistncePair(int[] nums,int k){
        Arrays.sort(nums);
        int lo=0,hi=nums[nums.length-1]-nums[0];
        while(lo<hi){
            int mi=(lo+hi)/2;
            int count=0,left=0;
            for(int right=0;right<nums.length;right++){
                while(nums[right]-nums[left]>mi) left++;
                count+=right-left;
            }
            if(count>=k) hi=mi;
            else lo=mi+1;
        }
        return lo;
    }
}
```

### [1254. 统计封闭岛屿的数目](https://leetcode-cn.com/problems/number-of-closed-islands/)

> 有一个二维矩阵 `grid` ，每个位置要么是陆地（记号为 `0` ）要么是水域（记号为 `1` ）。
>
> 我们从一块陆地出发，每次可以往上下左右 4 个方向相邻区域走，能走到的所有陆地区域，我们将其称为一座「**岛屿**」。
>
> 如果一座岛屿 **完全** 由水域包围，即陆地边缘上下左右所有相邻区域都是水域，那么我们将其称为 「**封闭岛屿**」。
>
> 请返回封闭岛屿的数目。

这是一道深度搜索的题目。

```java
class Solution{
    private int val;
    public int closedIsland(int[][] grid){
        int closedLandNum=0;
        for(int i=0;i<grid.length;i++){
            for(int j=0;j<grid[0].length;j++){
                if(grid[i][j]==0){
                    val=1;
                    dfs(grid,i,j);
                    closedLandNum+=val;
                }
            }
        }
        return closedLandNum;
    }
    public void dfs(int[][] grid,int i,int j){
        if(i<0||i>=grid.length||j<0||j>=grid[0].length){
            val=0;
            return ;
        }
        if(grid[i][j]!=0) return ;
        grid[i][j]=1;
        dfs(grid,i+1,j);
        dfs(grid,i-1,j);
        dfs(grid,i,j+1);
        dfs(grid,i,j-1);
    }
}
```

### [433. 最小基因变化](https://leetcode-cn.com/problems/minimum-genetic-mutation/)

>一条基因序列由一个带有8个字符的字符串表示，其中每个字符都属于 `"A"`, `"C"`, `"G"`, `"T"`中的任意一个。
>
>假设我们要调查一个基因序列的变化。**一次**基因变化意味着这个基因序列中的**一个**字符发生了变化。
>
>例如，基因序列由`"AACCGGTT"` 变化至 `"AACCGGTA" `即发生了一次基因变化。
>
>与此同时，每一次基因变化的结果，都需要是一个合法的基因串，即该结果属于一个基因库。
>
>现在给定3个参数 — start, end, bank，分别代表起始基因序列，目标基因序列及基因库，请找出能够使起始基因序列变化为目标基因序列所需的最少变化次数。如果无法实现目标变化，请返回 -1。
>
>**注意：**
>
>1. 起始基因序列默认是合法的，但是它并不一定会出现在基因库中。
>2. 如果一个起始基因序列需要多次变化，那么它每一次变化之后的基因序列都必须是合法的。
>3. 假定起始基因序列与目标基因序列是不一样的。

这是一道BFS，不过为了加快搜索的速度可以使用双端BFS。

```java
public int minMutation(String start, String end, String[] bank) {
        LinkedList<String> startQueue = new LinkedList<>();
        LinkedList<String> endQueue = new LinkedList<>();

        startQueue.add(start);
        endQueue.add(end);

        int res = 0;
        Set<String> dict = new HashSet<>();
        Set<String> mem = new HashSet<>();
        for (String b : bank) {
            dict.add(b);
        }
        if (!dict.contains(end)) {
            return -1;
        }
        while (!startQueue.isEmpty() && !endQueue.isEmpty()) {
            res++;

            LinkedList<String> tmpResQueue = new LinkedList<>();
            if (startQueue.size() > endQueue.size()) {
                LinkedList tmp = startQueue;
                startQueue = endQueue;
                endQueue = tmp;
            }
            int size = startQueue.size();
            for (int i = 0; i < size; i++) {
                String cur = startQueue.poll();
                for (int j = 0; j < bank.length; j++) {
                    if (vertify(cur, bank[j])) {
                        if (endQueue.contains(bank[j])) {
                            return res;
                        }
                        if (!mem.contains(bank[j])) {
                            tmpResQueue.add(bank[j]);
                            mem.add(bank[j]);
                        }
                    }
                }
            }
            if (startQueue.size() <= endQueue.size()) {
                startQueue = tmpResQueue;
            } else {
                endQueue = tmpResQueue;
            }
        }
        return -1;
    }

    public boolean vertify(String str1, String str2) {
        int count = 0;
        for (int i = 0; i < str1.length(); i++) {
            if (str1.charAt(i) != str2.charAt(i)) {
                count++;
            }
        }
        if (count > 1) {
            return false;
        }
        return true;
    }
```

### [1020. 飞地的数量](https://leetcode-cn.com/problems/number-of-enclaves/)

> 给出一个二维数组 `A`，每个单元格为 0（代表海）或 1（代表陆地）。
>
> 移动是指在陆地上从一个地方走到另一个地方（朝四个方向之一）或离开网格的边界。
>
> 返回网格中**无法**在任意次数的移动中离开网格边界的陆地单元格的数量。

这一题是DFS，在DFS的过程中将1改成0，剩下的1就是答案。

```java
class Solution{
    public int numEnclaves(int[][] grid){
        if(grid==null||grid.length==0||grid[0].length==0){
            return 0;
        }
        int r=grid.length,c=grid[0].length;
        for(int i=0;i<r;i++){
            dfs(grid,i,0);
            dfs(grid,i,c-1);
        }
        for(int i=0;i<c;i++){
            dfs(grid,0,i);
            dfs(grid,r-1,i);
        }
        
        int res=0;
        for(int i=0;i<r;i++){
            for(int j=0;j<c;j++){
                if(grid[i][j]==1){
                    res++;
                }
            }
        }
        return res;
    }
    
    public void dfs(int[][] grid,int r,int c){
        if(r<0||r>=grid.length||c<0||c>=grid[0].length){
            return ;
        }
        if(grid[r][c]==0) return ;
        
        grid[r][c]=0;
        dfs(grid,r+1,c);
        dfs(grid,r-1,c);
        dfs(grid,r,c+1);
        dfs()
    }
}
```

### [1091. 二进制矩阵中的最短路径](https://leetcode-cn.com/problems/shortest-path-in-binary-matrix/)

>给你一个 n x n 的二进制矩阵 grid 中，返回矩阵中最短 畅通路径 的长度。如果不存在这样的路径，返回 -1 。
>
>二进制矩阵中的 畅通路径 是一条从 左上角 单元格（即，(0, 0)）到 右下角 单元格（即，(n - 1, n - 1)）的路径，该路径同时满足下述要求：
>
>路径途经的所有单元格都的值都是 0 。
>路径中所有相邻的单元格应当在 8 个方向之一 上连通（即，相邻两单元之间彼此不同且共享一条边或者一个角）。
>畅通路径的长度 是该路径途经的单元格总数

最短路径问题，可以使用BFS来解决

```java
class Solution {
    private class Point{
        public int x;
        public int y;
        public Point(int x,int y){
            this.x=x;
            this.y=y;
        }

    }
    public int shortestPathBinaryMatrix(int[][] grid) {
        LinkedList<Point> queue=new LinkedList<>();
        int[][] dict=new int[][]{{1, -1}, {1, 0}, {1, 1}, {0, 1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}};
        if(grid==null||grid.length==0||grid[0].length==0){
            return -1;
        }
        if(grid[0][0]==1||grid[grid.length-1][grid[0].length-1]==1){
            return -1;
        }
        int res=0;
        queue.add(new Point(0,0));
        grid[0][0]=1;
        while(!queue.isEmpty()){
            res++;
            int size=queue.size();
            for(int i=0;i<size;i++){
                Point cur=queue.poll();
                int x=cur.x;
                int y=cur.y;
                if(x==grid.length-1&&y==grid[0].length-1){
                    return res;
                }
                
                for(int k=0;k<8;k++){
                    int x1=x+dict[k][0];
                    int y1=y+dict[k][1];
                    if(x1>=0&&x1<grid.length&&y1>=0&&y1<grid[0].length&&grid[x1][y1]==0){
                        grid[x1][y1]=1;
                        queue.add(new Point(x1,y1));
                    }
                }
            }
        }
        return -1;

    }
}
```

### [165. 比较版本号](https://leetcode-cn.com/problems/compare-version-numbers/)

> 给你两个版本号 `version1` 和 `version2` ，请你比较它们。
>
> 版本号由一个或多个修订号组成，各修订号由一个 `'.'` 连接。每个修订号由 **多位数字** 组成，可能包含 **前导零** 。每个版本号至少包含一个字符。修订号从左到右编号，下标从 0 开始，最左边的修订号下标为 0 ，下一个修订号下标为 1 ，以此类推。例如，`2.5.33` 和 `0.1` 都是有效的版本号。
>
> 比较版本号时，请按从左到右的顺序依次比较它们的修订号。比较修订号时，只需比较 **忽略任何前导零后的整数值** 。也就是说，修订号 `1` 和修订号 `001` **相等** 。如果版本号没有指定某个下标处的修订号，则该修订号视为 `0` 。例如，版本 `1.0` 小于版本 `1.1` ，因为它们下标为 `0` 的修订号相同，而下标为 `1` 的修订号分别为 `0` 和 `1` ，`0 < 1` 。
>
> 返回规则如下：
>
> - 如果 `*version1* > *version2*` 返回 `1`，
> - 如果 `*version1* < *version2*` 返回 `-1`，
> - 除此之外返回 `0`。

字符串处理问题，就两个version通过.进行分隔开，然后进行比对。

```go
func compareVersion(version1 string, version2 string) int {
    v1:=decode(version1)
    v2:=decode(version2)
    for len(v1)<len(v2){
        v1=append(v1,0)
    }
    for len(v2)<len(v1){
        v2=append(v2,0)
    }

    for i:=0;i<len(v1);i++{
        if v1[i]!=v2[i]{
            if v1[i]>v2[i]{
                return 1
            }else{
                return -1
            }
        }
    }
    return 0
}

func decode(v string)[]int{
    strs:=strings.Split(v,".")
    out:=make([]int,0)
    for _,str:=range strs{
        n,_:=strconv.Atoi(str)
        out=append(out,n)
    }
    return out
}
```

### [剑指 Offer 57 - II. 和为s的连续正数序列](https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/)

> 输入一个正整数 `target` ，输出所有和为 `target` 的连续正整数序列（至少含有两个数）。
>
> 序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

这是一道滑动窗口的题目，首先把1-target的数字看成一个窗口。然后左右移动指针，使得指针内的sum和为target

```go
func findContinuousSequence(target int) [][]int {
    i,j,sum:=1,1,0
    res:=make([][]int,0)
    for i<target{
        if sum<target{
            sum+=j
            j++
        }else{
            if sum==target{
                tmp:=make([]int,0)
                for k:=i;k<j;k++{
                    tmp=append(tmp,k)
                }
                res=append(res,tmp)
            }
            sum-=i
            i++
        }
    }
    return res
}
```

### [1367. 二叉树中的列表](https://leetcode-cn.com/problems/linked-list-in-binary-tree/)

> 给你一棵以 `root` 为根的二叉树和一个 `head` 为第一个节点的链表。
>
> 如果在二叉树中，存在一条一直向下的路径，且每个点的数值恰好一一对应以 `head` 为首的链表中每个节点的值，那么请你返回 `True` ，否则返回 `False` 。
>
> 一直向下的路径的意思是：从树中某个节点开始，一直连续向下的路径。

这是一道递归题

```go
type ListNode struct{
    Val int
    Next *ListNode
}

type TreeNode struct{
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isSubPath(head *ListNode,root *TreeNode) bool{
    if root==nil{
        return false
    }
    return isSub(head,root)||isSubPath(head,root.Left)||isSubPath(head,root.Right);
}

func isSub(head *ListNode,root *TreeNode) bool{
    if head==nil{
        return true
    }
    if root==nil{
        return false
    }
    if head.Val!=root.Val{
        return false
    }
    return isSub(head.Next,root.Left)||isSub(head.Next,root.Right)
}
```

### [821. 字符的最短距离](https://leetcode-cn.com/problems/shortest-distance-to-a-character/)

> 给你一个字符串 `s` 和一个字符 `c` ，且 `c` 是 `s` 中出现过的字符。
>
> 返回一个整数数组 `answer` ，其中 `answer.length == s.length` 且 `answer[i]` 是 `s` 中从下标 `i` 到离它 **最近** 的字符 `c` 的 **距离** 。
>
> 两个下标 `i` 和 `j` 之间的 **距离** 为 `abs(i - j)` ，其中 `abs` 是绝对值函数。

解法：遍历两次字符串，从左到右更新一遍到c值得距离，在从右更新一遍到C值得距离【从左到右 假设c值在左边，从右到左 假设c值在右边】

```go
func shortestToChar(s string,c byte)[]int{
    n:=len(s)
    res:=make([]int,n)
    for i:=range res{
     	res[i]=n
    }
    
    for i:=0;i<n;i++{
        if s[i]==c{
            res[i]=0
        }else if i>0{
            res[i]=res[i-1]+1
        }
    }
    
    for i:=n-1;i>=0;i--{
        if i<n-1&&res[i+1]+1<res[i]{
            res[i]=res[i+1]+1
        }
    }
    return res
}
```

### [881. 救生艇](https://leetcode-cn.com/problems/boats-to-save-people/)

> 第 `i` 个人的体重为 `people[i]`，每艘船可以承载的最大重量为 `limit`。
>
> 每艘船最多可同时载两人，但条件是这些人的重量之和最多为 `limit`。
>
> 返回载到每一个人所需的最小船数。(保证每个人都能被船载)。

这是一道双指针贪心题，通过排序数组，让最轻(left)和最重的(right)一起坐船。

```go
func numRescueBoats(people []int, limit int)int{
    sort.Ints(people)
    left,right,res:=0,len(people)-1,0
    for left<=right{
        res++
        if people[left]+people[right]<=limit{
            left++
        }
        
    }
}
```

### [532. 数组中的 k-diff 数对](https://leetcode-cn.com/problems/k-diff-pairs-in-an-array/)

> 给定一个整数数组和一个整数 `**k**`，你需要在数组里找到 **不同的** k-diff 数对，并返回不同的 **k-diff 数对** 的数目。
>
> 这里将 **k-diff** 数对定义为一个整数对 `(nums[i], nums[j])`，并满足下述全部条件：
>
> - `0 <= i < j < nums.length`
> - `|nums[i] - nums[j]| == k`
>
> **注意**，`|val|` 表示 `val` 的绝对值。

看到这道题，我一开始想到的方法是双指针，但是这里数组的每个数字可以被多次使用，所以不好处理这种情况。

正确思路是通过map来查找某个数字，跟2sum有点像。

```go
func findPairs(nums []int,k int)int{
    if k<0||len(nums)==0{
        return 0
    }
    count:=0
    m:=make(map[int]int,len(nums))
    for _,n:=range nums{
        m[n]++
    }
    
    for key:=range m{
        if k==0&&m[key]>1{
            count++
            continue
        }
        if k>0&&m[key+k]>0{
            count++
        }
    }
    return count
}
```

### [599. 两个列表的最小索引总和](https://leetcode-cn.com/problems/minimum-index-sum-of-two-lists/)

> 假设Andy和Doris想在晚餐时选择一家餐厅，并且他们都有一个表示最喜爱餐厅的列表，每个餐厅的名字用字符串表示。
>
> 你需要帮助他们用**最少的索引和**找出他们**共同喜爱的餐厅**。 如果答案不止一个，则输出所有答案并且不考虑顺序。 你可以假设总是存在一个答案。

这里的最少索引和表示列表1和列表2相同的餐厅，它们的索引和。

这里可以使用map，map的key存储餐厅，value存储索引。

```go
func findRestaurant(list1 []string,list2 []string)[]string{
    m,ans:=make(map[string]int,0),make([]string,0)
    for index,item:=range list1{
        m[item]=index
    }
    
    for index,item:=range list2{
        if _,ok:=m[item];ok{
            m[item]+=index
            if len(ans)==0||m[item]==m[ans[0]]{
                ans=append(ans,item)
            }else if m[item]<m[ans[0]]{
                ans=[]string{item}
            }
        }
    }
    return ans
}
```

### [面试题 16.24. 数对和](https://leetcode-cn.com/problems/pairs-with-sum-lcci/)

> 设计一个算法，找出数组中两数之和为指定值的所有整数对。一个数只能属于一个数对。
>
> **示例 1:**
>
> ```
> 输入: nums = [5,6,5], target = 11
> 输出: [[5,6]]
> ```
>
> **示例 2:**
>
> ```
> 输入: nums = [5,6,5,6], target = 11
> 输出: [[5,6],[5,6]]
> ```
>
> **提示：**
>
> - `nums.length <= 100000`

通过排序，然后左右指针找到数对，由于一个数只能属于一个数对，即只能用一次，所以容易完成

```go
func pairSums(nums []int,target int)[][]int{
    ans:=make([][]int,0)
    sort.Ints(nums)
    left,right:=0,len(nums)-1
    for left<right{
        if nums[left]+nums[right]==target{
            ans=append(ans,[]int{nums[left],nums[right]})
            left++
            right--
        }else if nums[left]+nums[right]<target{
            left++
        }else{
            right--
        }
    }
    return ans
}
```

### [1742. 盒子中小球的最大数量](https://leetcode-cn.com/problems/maximum-number-of-balls-in-a-box/)

> 你在一家生产小球的玩具厂工作，有 `n` 个小球，编号从 `lowLimit` 开始，到 `highLimit` 结束（包括 `lowLimit` 和 `highLimit` ，即 `n == highLimit - lowLimit + 1`）。另有无限数量的盒子，编号从 `1` 到 `infinity` 。
>
> 你的工作是将每个小球放入盒子中，其中盒子的编号应当等于小球编号上每位数字的和。例如，编号 `321` 的小球应当放入编号 `3 + 2 + 1 = 6` 的盒子，而编号 `10` 的小球应当放入编号 `1 + 0 = 1` 的盒子。
>
> 给你两个整数 `lowLimit` 和 `highLimit` ，返回放有最多小球的盒子中的小球数量*。*如果有多个盒子都满足放有最多小球，只需返回其中任一盒子的小球数量。

遍历一遍，然后把结果放入map，取出最大值

```go
func countBalls(lowLimit int, highLimit int) int {
    m,ans:=make(map[int]int,0),0

    for i:=lowLimit;i<=highLimit;i++{
        t,j:=0,i
        for j>0{
            t+=j%10
            j/=10
        }
        if _,ok:=m[t];ok{
            m[t]+=1
        }else{
            m[t]=1
        }

        if ans<m[t]{
            ans=m[t]
        }
    }
    return ans
}
```

### [1487. 保证文件名唯一](https://leetcode-cn.com/problems/making-file-names-unique/)

> 给你一个长度为 `n` 的字符串数组 `names` 。你将会在文件系统中创建 `n` 个文件夹：在第 `i` 分钟，新建名为 `names[i]` 的文件夹。
>
> 由于两个文件 **不能** 共享相同的文件名，因此如果新建文件夹使用的文件名已经被占用，系统会以 `(k)` 的形式为新文件夹的文件名添加后缀，其中 `k` 是能保证文件名唯一的 **最小正整数** 。
>
> 返回长度为 *`n`* 的字符串数组，其中 `ans[i]` 是创建第 `i` 个文件夹时系统分配给该文件夹的实际名称。

map存储当前已经在res里的文件名了

```go 
func getFolderNames(names []string)[]string{
    m,res:=make(map[string]int,0),make([]string,0)
    for _,name:=range names{
        tmp:=name
        for _,ok:=m[tmp];ok;_,ok=m[tmp]{
            num:=m[name]
            m[name]+=1
            tmp=name+"("+strconv.Itoa(num)+")"
        }
        res=append(res,tmp)
        m[tmp]=1
        if _,ok:=m[name];!ok{
            m[name]=1
        }
    }
    return res
}
```



