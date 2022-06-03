package main

func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]bool{}
	res := make([]int, 0)
	for _, n := range nums1 {
		m[n] = true
	}
	for _, n := range nums2 {
		if _, ok := m[n]; ok {
			delete(m, n)
			res = append(res, n)
		}
	}
	return res
}

//func main() {
//	nums1 := []int{1, 2, 2, 1}
//	nums2 := []int{2, 2}
//	_ = intersection(nums1, nums2)
//}
