package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	h1, h2 := headA, headB
	for h1 != h2 {
		if h1 == nil {
			h1 = headB
		} else {
			h1 = h1.Next
		}
		if h2 == nil {
			h2 = headA
		} else {
			h2 = h2.Next
		}
	}
	return h1
}
