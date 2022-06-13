package main

func intersection(nums1 []int, nums2 []int) []int {
	m, res := make(map[int]bool), make([]int, 0)
	for _, v := range nums1 {
		m[v] = true
	}
	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			delete(m, v)
			res = append(res, v)
		}
	}
	return res
}

func main() {
	nums1, nums2 := []int{1, 2, 2, 1}, []int{2, 2}
	_ = intersection(nums1, nums2)
}
