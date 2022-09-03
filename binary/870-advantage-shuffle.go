package main

import "sort"

func advantageCount(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	sort.Ints(nums1)
	sorted2 := make([]int, n)
	for i := 0; i < n; i++ {
		sorted2[i] = i
	}
	sort.Slice(sorted2, func(i, j int) bool {
		return nums2[sorted2[i]] < nums2[sorted2[j]]
	})

	unless, res, i := make([]int, 0), make([]int, n), 0
	for _, index := range sorted2 {
		b := nums2[index]
		for i < n && nums1[i] <= b {
			unless = append(unless, nums1[i])
			i++
		}

		if i < n {
			res[index] = nums1[i]
			i++
		} else {
			res[index] = unless[0]
			unless = unless[1:]
		}
	}
	return res
}
