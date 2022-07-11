package main

func sortArray3(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return nums
	}
	end := len(nums) - 1
	for i := end / 2; i >= 0; i-- {
		heapify(nums, i, end)
	}

	for i := end; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		end--
		heapify(nums, 0, end)
	}
	return nums
}

func heapify(nums []int, root, end int) {
	child := root*2 + 1
	if child > end {
		return
	}
	if child < end && nums[child] < nums[child+1] {
		child += 1
	}
	if nums[root] > nums[child] {
		return
	}
	nums[root], nums[child] = nums[child], nums[root]
	heapify(nums, child, end)
}

//func main() {
//	nums := []int{5, 2, 3, 1}
//	fmt.Printf("%v \n", sortArray3(nums))
//}
