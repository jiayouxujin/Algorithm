package main

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			fast = head
			for slow != fast {
				fast = fast.Next
				slow = slow.Next
			}
			return fast
		}
	}
	return nil
}
