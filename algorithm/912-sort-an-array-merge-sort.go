package main

import "fmt"

func sortArray2(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return nums
	}
	mergeSort(nums, 0, len(nums)-1)
	return nums
}

func mergeSort(nums []int, l, r int) {
	if l < r {
		mid := l + (r-l)/2
		mergeSort(nums, l, mid)
		mergeSort(nums, mid+1, r)
		merge(nums, l, mid, r)
	}
}

func merge(nums []int, l, mid, r int) {
	tmp := make([]int, 0)
	i, j := l, mid+1
	for i <= mid && j <= r {
		if nums[i] <= nums[j] {
			tmp = append(tmp, nums[i])
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}
	for i <= mid {
		tmp = append(tmp, nums[i])
		i++
	}
	for j <= r {
		tmp = append(tmp, nums[j])
		j++
	}

	for k := l; k <= r; k++ {
		nums[k] = tmp[k-l]
	}
}

func main() {
	nums := []int{5, 2, 3, 1}
	fmt.Printf("%v \n", sortArray2(nums))
}
