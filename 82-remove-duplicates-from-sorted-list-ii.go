package main

func deleteDuplicates(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	p := dummyHead

	for p.Next != nil {
		q := p.Next.Next
		for q != nil && q.Val == p.Next.Val {
			q = q.Next
		}
		if p.Next.Next == q {
			p = p.Next
		} else {
			p.Next = q
		}
	}
	return dummyHead.Next
}
