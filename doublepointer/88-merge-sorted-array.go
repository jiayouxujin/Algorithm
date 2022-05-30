package main

//给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
//
//请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
func merge(nums1 []int, m int, nums2 []int, n int) {
	len1, len2, length := m-1, n-1, len(nums1)-1
	for len1 >= 0 && len2 >= 0 {
		if nums1[len1] > nums2[len2] {
			nums1[length] = nums1[len1]
			length--
			len1--
		} else {
			nums1[length] = nums2[len2]
			length--
			len2--
		}
	}

	for len2 >= 0 {
		nums1[length] = nums2[len2]
		length--
		len2--
	}
}

//func main() {
//	nums1, m, nums2, n := []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3
//	merge(nums1, m, nums2, n)
//}
