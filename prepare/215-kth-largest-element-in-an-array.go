package main

import "math/rand"

func findKthLargest(nums []int, k int) int {
	k = len(nums) - k + 1
	return quickSortVar(nums, k, 0, len(nums)-1)
}

func quickSortVar(nums []int, target, left, right int) int {
	if left >= right {
		return nums[left]
	}
	p := partition(nums, left, right)
	q := p - left + 1
	if q > target {
		return quickSortVar(nums, target, left, p-1)
	} else if q < target {
		return quickSortVar(nums, target-q, p+1, right)
	}
	return nums[p]
}

func partition2(nums []int, left, right int) int {
	p := rand.Intn(right-left+1) + left
	nums[right], nums[p] = nums[p], nums[right]
	i := left - 1
	for j := left; j < right; j++ {
		if nums[j] <= nums[right] {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[right] = nums[right], nums[i+1]
	return i + 1
}
