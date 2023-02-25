package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// version 1
//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	if l1 == nil {
//		return l2
//	}
//	if l2 == nil {
//		return l1
//	}
//	head, prev, cur := &ListNode{Next: nil}, &ListNode{Next: nil}, &ListNode{Next: nil}
//	for l1 != nil && l2 != nil {
//		if l1.Val <= l2.Val {
//			cur = l1
//			l1 = l1.Next
//		} else {
//			cur = l2
//			l2 = l2.Next
//		}
//		if head.Next == nil {
//			head.Next = cur
//		}
//		prev.Next = cur
//		prev = cur
//	}
//	if l1 != nil {
//		prev.Next = l1
//	} else {
//		prev.Next = l2
//	}
//	return head.Next
//}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}
