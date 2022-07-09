package main

import (
	"math/rand"
)

func sortArray(nums []int) []int {
	if nums == nil || len(nums) <= 1 {
		return nums
	}

	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	q := quickPartition(nums, left, right)
	quickSort(nums, left, q-1)
	quickSort(nums, q+1, right)
}

func quickPartition(nums []int, left, right int) int {
	q := left + rand.Intn(right-left+1)
	nums[q], nums[right] = nums[right], nums[q]
	i := left - 1
	for j := left; j < right; j++ {
		if nums[j] < nums[right] {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[right] = nums[right], nums[i+1]
	return i + 1
}

//func main() {
//	nums := []int{5, 2, 3, 1}
//	nums = sortArray(nums)
//	fmt.Printf("%v \n", nums)
//}
