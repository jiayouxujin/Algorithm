package main

import "strconv"

func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	numStr := toStringArray(nums)
	quickSortString(numStr, 0, len(numStr)-1)
	res := ""
	for _, str := range numStr {
		if res == "0" && str == "0" {
			continue
		}
		res = res + str
	}
	return res
}

func toStringArray(nums []int) []string {
	numStr := make([]string, 0)
	for _, num := range nums {
		numStr = append(numStr, strconv.Itoa(num))
	}
	return numStr
}

func partitionString(a []string, l, r int) int {
	pivot := a[r]
	i := l - 1
	for j := l; j < r; j++ {
		ajStr := a[j] + pivot
		pivotStr := pivot + a[j]
		if ajStr > pivotStr {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1

}

func quickSortString(a []string, l, r int) {
	if l >= r {
		return
	}
	p := partitionString(a, l, r)
	quickSortString(a, l, p-1)
	quickSortString(a, p+1, r)
}
