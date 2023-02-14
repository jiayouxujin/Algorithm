package main

func sortArray3(nums []int) []int {
	n := len(nums)
	for i := n / 2; i >= 0; i-- {
		heapify(nums, n, i)
	}

	for i := n - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, i, 0)
	}
	return nums
}

func heapify(nums []int, n, i int) {
	largest := i
	left, right := 2*i, 2*i+1
	if left < n && nums[left] > nums[largest] {
		largest = left
	}
	if right < n && nums[right] > nums[largest] {
		largest = right
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		heapify(nums, n, largest)
	}
}
