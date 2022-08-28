package main

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyHead := &ListNode{Next: head}
	p := dummyHead
	for p.Next != nil {
		q := p.Next.Next
		for q != nil && p.Next.Val == q.Val {
			q = q.Next
		}
		if q == p.Next.Next {
			p = p.Next
		} else {
			p.Next = q
		}
	}
	return dummyHead.Next
}
