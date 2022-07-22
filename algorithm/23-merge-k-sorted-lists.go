package main

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeList(lists, 0, len(lists)-1)
}

func mergeList(lists []*ListNode, left, right int) *ListNode {
	if left >= right {
		return lists[left]
	}
	mid := left + (right-left)/2
	l := mergeList(lists, left, mid)
	r := mergeList(lists, mid+1, right)
	return mergeTwo(l, r)
}

func mergeTwo(a, b *ListNode) *ListNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.Val <= b.Val {
		a.Next = mergeTwo(a.Next, b)
		return a
	}
	b.Next = mergeTwo(a, b.Next)
	return b
}
