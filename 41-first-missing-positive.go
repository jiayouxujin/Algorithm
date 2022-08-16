package main

func firstMissingPositive(nums []int) int {
	m := make(map[int]struct{})
	for _, v := range nums {
		m[v] = struct{}{}
	}
	for i := 1; i <= len(nums); i++ {
		if _, ok := m[i]; ok {
			continue
		}
		return i
	}
	return len(nums) + 1
}
