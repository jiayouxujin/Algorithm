package main

func sortArray2(nums []int) []int {
	mergeSort(nums, 0, len(nums)-1)
	return nums
}

func mergeSort(nums []int, left, right int) {
	if left < right {
		mid := left + (right-left)/2
		mergeSort(nums, left, mid)
		mergeSort(nums, mid+1, right)
		merge(nums, left, mid, right)
	}
}

func merge(nums []int, left, mid, right int) {
	tmp := make([]int, 0)
	i, j := left, mid+1
	for i <= mid && j <= right {
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
	for j <= right {
		tmp = append(tmp, nums[j])
		j++
	}

	for k := left; k <= right; k++ {
		nums[k] = tmp[k-left]
	}
}
