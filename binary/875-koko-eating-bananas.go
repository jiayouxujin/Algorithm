package main

func minEatingSpeed(piles []int, h int) int {
	lo, hi := 1, 1000000001
	for lo < hi {
		mid := lo + (hi-lo)/2
		n := f(piles, mid)
		if n == h {
			hi = mid
		} else if n > h {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func f(piles []int, x int) int {
	hours := 0
	for i := 0; i < len(piles); i++ {
		hours += piles[i] / x
		if piles[x]%x > 0 {
			hours++
		}
	}
	return hours
}
