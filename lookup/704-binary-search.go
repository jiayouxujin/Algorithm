package main

func search2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}

//func main() {
//	nums, target := []int{-1, 0, 3, 5, 9, 12}, 9
//	fmt.Printf("%v \n", search2(nums, target))
//}
