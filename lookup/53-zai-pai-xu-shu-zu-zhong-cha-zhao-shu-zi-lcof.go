package main

func search3(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	return rightBound(nums, target) - leftBound(nums, target) + 1
}

func rightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}

func leftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left >= len(nums) || nums[left] != target {
		return 0
	}
	return left
}

//func main() {
//	nums, target := []int{5, 7, 7, 8, 8, 10}, 8
//	fmt.Printf("%v \n", search3(nums, target))
//}
