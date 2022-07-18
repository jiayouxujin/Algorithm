package main

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeK(lists, 0, len(lists)-1)
}

func mergeK(lists []*ListNode, left, right int) *ListNode {
	if left >= right {
		return lists[left]
	}
	mid := left + (right-left)/2
	a := mergeK(lists, left, mid)
	b := mergeK(lists, mid+1, right)
	return mergeTwoListss(a, b)
}

func mergeTwoListss(a, b *ListNode) *ListNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.Val <= b.Val {
		a.Next = mergeTwoLists(a.Next, b)
		return a
	}
	b.Next = mergeTwoLists(a, b.Next)
	return b
}
