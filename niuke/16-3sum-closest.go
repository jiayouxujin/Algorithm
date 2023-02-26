package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	diff, res := math.MaxInt, 0
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return target
			} else if sum > target {
				if sum-target < diff {
					diff = sum - target
					res = sum
				}
				right--
			} else {
				if target-sum < diff {
					diff = target - sum
					res = sum
				}
				left++
			}
		}
	}
	return res
}
