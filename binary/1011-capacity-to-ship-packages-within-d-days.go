package main

func shipWithinDays(weights []int, days int) int {
	lo, hi := getMax(weights), getSum(weights)
	for lo < hi {
		mid := lo + (hi-lo)/2
		n := splitArray(weights, mid)
		if n == days {
			hi = mid
		} else if n > days {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
