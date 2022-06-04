package main

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		if target == nums[mid] {
			return mid
		}
		if nums[left] <= nums[mid] {
			if target < nums[mid] && nums[left] <= target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	if nums[left] == target {
		return left
	} else {
		return -1
	}
}

//func main() {
//	nums, target := []int{4, 5, 6, 7, 0, 1, 2}, 0
//	fmt.Printf("%v \n", search(nums, target))
//}
