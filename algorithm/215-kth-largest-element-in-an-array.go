package main

import (
	"math/rand"
)

func findKthLargest(nums []int, k int) int {
	if nums == nil || len(nums) < k {
		return 0
	}
	k = len(nums) - k + 1
	return quickFind(nums, 0, len(nums)-1, k)
}

func quickFind(nums []int, i, j, k int) int {
	if i >= j {
		return nums[i]
	}
	p := quickPartition(nums, i, j)
	q := p - i + 1
	if q == k {
		return nums[p]
	}

	if k < q {
		return quickFind(nums, i, p-1, k)
	} else {
		return quickFind(nums, p+1, j, k-q)
	}
}

func quickPartition(nums []int, l, r int) int {
	p := rand.Intn(r-l+1) + l
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
//	nums, k := []int{3, 2, 1, 5, 6, 4}, 2
//	fmt.Printf("%v \n", findKthLargest(nums, k))
//}
