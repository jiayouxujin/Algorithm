package main

import "math/rand"

type Solution struct {
	preSum []int
}

func Constructor(w []int) Solution {
	preSum := make([]int, len(w))
	for i := 0; i < len(w); i++ {
		if i == 0 {
			preSum[i] = w[i]
		} else {
			preSum[i] = preSum[i-1] + w[i]
		}
	}
	return Solution{preSum: preSum}
}

func (this *Solution) PickIndex() int {
	n := rand.Intn(this.preSum[len(this.preSum)-1]) + 1
	lo, hi := 0, len(this.preSum)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if this.preSum[mid] == n {
			hi = mid
		} else if this.preSum[mid] > n {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
