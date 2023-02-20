package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || right < left {
		return head
	}
	newHead := &ListNode{Next: head}
	prev := newHead
	for count := 0; prev != nil && count < left-1; count++ {
		prev = prev.Next
	}
	if prev == nil {
		return head
	}
	cur := prev.Next
	for i := 0; i < right-left; i++ {
		tmp := prev.Next
		prev.Next = cur.Next
		cur.Next = cur.Next.Next
		prev.Next.Next = tmp
	}
	return newHead.Next
}
