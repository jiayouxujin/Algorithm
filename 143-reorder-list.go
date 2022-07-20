package main

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	//找到中点
	p1, p2 := head, head
	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}

	//反转后半部分
	preMiddle, preCurrent := p1, p1.Next
	for preCurrent.Next != nil {
		next := preCurrent.Next
		preCurrent.Next = next.Next
		next.Next = preMiddle.Next
		preMiddle.Next = next
	}

	//重新拼接链表
	p1, p2 = head, preMiddle.Next
	for p1 != preMiddle {
		preMiddle.Next = p2.Next
		p2.Next = p1.Next
		p1.Next = p2
		p1 = p2.Next
		p2 = preMiddle.Next
	}
}
