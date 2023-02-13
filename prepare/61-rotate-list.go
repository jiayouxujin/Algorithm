package main

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	newHead, len := &ListNode{Next: head}, 0
	cur := newHead
	for cur.Next != nil {
		cur = cur.Next
		len++
	}

	if k%len == 0 {
		return head
	}

	cur.Next = head
	cur = newHead
	for i := len - k%len; i > 0; i-- {
		cur = cur.Next
	}
	res := &ListNode{Next: cur.Next}
	cur.Next = nil
	return res.Next
}
