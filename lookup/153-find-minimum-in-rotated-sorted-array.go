package main

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}

//func main() {
//	nums := []int{3, 4, 5, 1, 2}
//	fmt.Printf("%v \n", findMin(nums))
//}
