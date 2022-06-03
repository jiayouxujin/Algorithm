package main

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	delta := int(^uint(0) >> 1)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		sum := nums[i] + twoSumClosest(nums, i+1, target-nums[i])
		if abs(delta) > abs(target-sum) {
			delta = target - sum
		}
	}
	return target - delta
}

func twoSumClosest(nums []int, start, target int) int {
	lo, hi := start, len(nums)-1
	delta := int(^uint(0) >> 1)
	for lo < hi {
		sum := nums[lo] + nums[hi]
		if abs(delta) > abs(target-sum) {
			delta = target - sum
		}
		if sum < target {
			lo++
		} else {
			hi--
		}
	}
	return target - delta
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
//func main() {
//	nums, target := []int{-1, 2, 1, -4}, 1
//	fmt.Printf("%v \n", threeSumClosest(nums, target))
//}
