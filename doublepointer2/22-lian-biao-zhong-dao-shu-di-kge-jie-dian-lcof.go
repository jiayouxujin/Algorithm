package main

func getKthFromEnd(head *ListNode, k int) *ListNode {
	slow, fast := head, head
	for fast != nil {
		if k <= 0 {
			slow = slow.Next
		}
		fast = fast.Next
		k--
	}
	return slow
}
