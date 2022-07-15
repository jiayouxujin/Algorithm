package main

func sortArray22(nums []int) []int {
	mergeSort2(nums, 0, len(nums)-1)
	return nums
}

func mergeSort2(nums []int, left, right int) {
	if left < right {
		mid := left + (right-left)/2
		mergeSort2(nums, left, mid)
		mergeSort2(nums, mid+1, right)
		merge2(nums, left, mid, right)
	}
}

func merge2(nums []int, left, mid, right int) {
	i, j := left, mid+1
	tmp := make([]int, 0)
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
	return
}

//func main() {
//	nums := []int{5, 2, 3, 1}
//	fmt.Printf("%v \n", sortArray22(nums))
//}
