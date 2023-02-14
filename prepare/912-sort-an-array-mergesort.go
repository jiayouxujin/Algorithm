package main

func sortArray(nums []int) []int {
	mergeSort(nums, 0, len(nums)-1)
	return nums
}

func mergeSort(nums []int, left, right int) {
	if left < right {
		mid := (right-left)/2 + left
		mergeSort(nums, left, mid)
		mergeSort(nums, mid+1, right)
		merge(nums, left, mid, right)
	}
}

func merge(nums []int, l, m, r int) {
	temp := make([]int, 0)
	i, j := l, m+1
	for i <= m && j <= r {
		if nums[i] <= nums[j] {
			temp = append(temp, nums[i])
			i++
		} else {
			temp = append(temp, nums[j])
			j++
		}
	}

	for i <= m {
		temp = append(temp, nums[i])
		i++
	}

	for j <= r {
		temp = append(temp, nums[j])
		j++
	}

	for i := l; i <= r; i++ {
		nums[i] = temp[i-l]
	}
}
