package main

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	newHead := &ListNode{Val: 0, Next: head}
	len, cur := 0, newHead
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
	res := &ListNode{Val: 0, Next: cur.Next}
	cur.Next = nil
	return res.Next
}
