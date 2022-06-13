package main

func findDuplicate(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		count := 0
		for _, v := range nums {
			if v <= mid {
				count++
			}
		}
		if count > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

//func main() {
//	nums := []int{1, 3, 4, 2, 2}
//	fmt.Printf("%v \n", findDuplicate(nums))
//}
