/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	length:=len(lists)
	if length<1{
		return nil
	}

	if length==1{
		return lists[0]
	}

	num:=length/2
	left:=mergeKLists(lists[:num])
	right:=mergeKLists(lists[num:])
	return mergeTowLists(left,right)
}

func mergeTowLists(l1 *ListNode,l2 *ListNode)  *ListNode{
	if l1==nil{
		return l2
	}
	if l2==nil{
		return l1
	}
	if l1.Val<l2.Val{
		l1.Next=mergeTowLists(l1.Next,l2)
		return l1
	}
	l2.Next=mergeTowLists(l1,l2.Next)
	return l2
}
// @lc code=end

