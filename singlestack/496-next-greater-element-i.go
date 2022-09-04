package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m, stk, res := make(map[int]int), make([]int, 0), make([]int, 0)
	//先计算出nums2的单调栈
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stk) > 0 && stk[len(stk)-1] < nums2[i] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			m[nums2[i]] = -1
		} else {
			m[nums2[i]] = stk[len(stk)-1]
		}
		stk = append(stk, nums2[i])
	}

	for i := 0; i < len(nums1); i++ {
		res = append(res, m[nums1[i]])
	}
	return res
}
