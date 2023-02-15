package main

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res, start, prev := make([][]int, 0), intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= prev {
			prev = max(prev, intervals[i][1])
		} else {
			res = append(res, []int{start, prev})
			start = intervals[i][0]
			prev = intervals[i][1]
		}
	}
	res = append(res, []int{start, prev})
	return res
}
