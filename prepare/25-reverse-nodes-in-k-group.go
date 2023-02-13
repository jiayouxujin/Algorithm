package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	a, b := head, head
	for i := 0; i < k; i++ {
		if b != nil {
			b = b.Next
		} else {
			return a
		}
	}
	newHead := reverseList(a, b)
	a.Next = reverseKGroup(b, k)
	return newHead
}

func reverseList(a, b *ListNode) *ListNode {
	pre, cur := a, a.Next
	for cur != b && cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
