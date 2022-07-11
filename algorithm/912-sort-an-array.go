package main

import (
	"math/rand"
)

func sortArray(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return nums
	}
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	p := partition(nums, l, r)
	quickSort(nums, l, p-1)
	quickSort(nums, p+1, r)
	return
}

func partition(nums []int, l, r int) int {
	p := l + rand.Intn(r-l+1)
	nums[p], nums[r] = nums[r], nums[p]
	i := l - 1
	for j := l; j < r; j++ {
		if nums[j] <= nums[r] {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}

//func main() {
//	nums := []int{5, 2, 3, 1}
//	fmt.Printf("%v \n", sortArray(nums))
//}
