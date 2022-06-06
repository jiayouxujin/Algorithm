package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	len := len(nums1) - 1
	m -= 1
	n -= 1
	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[len] = nums1[m]
			len--
			m--
		} else {
			nums1[len] = nums2[n]
			len--
			n--
		}
	}
	for n >= 0 {
		nums1[len] = nums2[n]
		len--
		n--
	}
}

//func main() {
//	nums1, m, nums2, n := []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3
//	merge(nums1, m, nums2, n)
//}
