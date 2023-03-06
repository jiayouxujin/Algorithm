package main

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	//找到中点
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	//划分左右两个部分
	left, right := head, slow.Next
	slow.Next = nil
	//反转后半部分
	var prev *ListNode
	for right != nil {
		next := right.Next
		right.Next = prev
		prev = right
		right = next
	}

	//拼接 left和prev
	for prev != nil {
		next := left.Next
		left.Next = prev
		prev = prev.Next
		left.Next.Next = next
		left = next
	}
}
