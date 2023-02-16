package main

func intersection(nums1 []int, nums2 []int) []int {
	res, m := make([]int, 0), make(map[int]int)
	for _, n := range nums1 {
		m[n]++
	}
	for _, n := range nums2 {
		if _, ok := m[n]; ok {
			delete(m, n)
			res = append(res, n)
		}
	}
	return res
}
