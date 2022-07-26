package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2
	dummy := &ListNode{Val: -1}
	p := dummy
	carry := 0

	for p1 != nil || p2 != nil || carry > 0 {
		val := carry
		if p1 != nil {
			val += p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			val += p2.Val
			p2 = p2.Next
		}
		carry = val / 10
		val = val % 10
		p.Next = &ListNode{Val: val}
		p = p.Next
	}
	return dummy.Next
}
