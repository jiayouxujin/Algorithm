package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || right <= left {
		return head
	}
	newHead := &ListNode{Val: 0, Next: head}
	prev := newHead
	for count := 0; prev.Next != nil && count < left-1; count++ {
		prev = prev.Next
	}

	if prev.Next == nil {
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
