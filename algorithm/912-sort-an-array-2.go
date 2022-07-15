package main

import (
	"math/rand"
)

//快排
func sortArray21(nums []int) []int {
	quickSort2(nums, 0, len(nums)-1)
	return nums
}

func quickSort2(nums []int, left, right int) {
	if left >= right {
		return
	}

	p := partition2(nums, left, right)
	quickSort2(nums, left, p-1)
	quickSort2(nums, p+1, right)
}

func partition2(nums []int, left, right int) int {
	p := rand.Intn(right-left+1) + left
	nums[p], nums[right] = nums[right], nums[p]

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

//func main() {
//	nums := []int{5, 2, 3, 1}
//	fmt.Printf("%v \n", sortArray21(nums))
//}
