package main

func sortArray3(nums []int) []int {
	//最后一个非叶子节点
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
	//max child
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
//	nums = sortArray3(nums)
//	fmt.Printf("%v \n", nums)
//}
