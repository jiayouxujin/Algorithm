package main

//堆排序
func sortArray23(nums []int) []int {
	end := len(nums) - 1
	for i := end / 2; i >= 0; i-- {
		heapify2(nums, i, end)
	}
	for i := end; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		end--
		heapify2(nums, 0, end)
	}
	return nums
}

func heapify2(nums []int, root, end int) {
	child := root*2 + 1
	if child > end {
		return
	}
	if child < end && nums[child] < nums[child+1] {
		child = child + 1
	}
	if nums[root] >= nums[child] {
		return
	}
	nums[root], nums[child] = nums[child], nums[root]
	root = child
	heapify2(nums, root, end)
}

//func main() {
//	nums := []int{5, 2, 3, 1}
//	fmt.Printf("%v \n", sortArray22(nums))
//}
