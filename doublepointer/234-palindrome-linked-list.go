package main

var left *ListNode

//version 1.0
//func isPalindrome(head *ListNode) bool {
//	left=head
//	return traverse(head);
//}
//
//func traverse(right *ListNode) bool {
//	if right==nil{
//		return true
//	}
//	res:=traverse(right.Next)
//	res=res&&(left.Val==right.Val)
//	left=left.Next
//	return res
//}

//version 2.0
func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	left, right := head, reverse(slow)
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
