package main

func merge2(nums1 []int, m int, nums2 []int, n int) {
	l := len(nums1) - 1
	m--
	n--
	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[l] = nums1[m]
			m--
		} else {
			nums1[l] = nums2[n]
			n--
		}
		l--
	}
	if m >= 0 {
		return
	} else {
		for n >= 0 {
			nums1[l] = nums2[n]
			n--
			l--
		}
	}
}
