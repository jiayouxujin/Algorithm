package main

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
	newHead := reverseRange(a, b)
	a.Next = reverseKGroup(b, k)
	return newHead
}

func reverseRange(a, b *ListNode) *ListNode {
	cur, prev := a, &ListNode{}
	for cur != nil && cur != b {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return prev
}
