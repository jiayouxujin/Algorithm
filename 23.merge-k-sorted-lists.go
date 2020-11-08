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

func mergeTowLists(left,right *ListNode)  *ListNode{
	if left==nil{
		return right
	}
	if right==nil{
		return left
	}

	if left.Val<right.Val{
		left.Next=mergeTowLists(left.Next,right)
		return left
	}
	right.Next=mergeTowLists(left,right.Next)
	return right
}
// @lc code=end

