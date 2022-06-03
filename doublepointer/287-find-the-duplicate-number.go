package main

func findDuplicate(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid, count := low+(high-low)>>1, 0
		for _, v := range nums {
			if v <= mid {
				count++
			}
		}
		if count > mid {
			high = mid
		} else {
			low = mid+1
		}
	}
	return low
}

//func main() {
//	nums := []int{1, 3, 4, 2, 2}
//	fmt.Printf("%v \n", findDuplicate(nums))
//}
