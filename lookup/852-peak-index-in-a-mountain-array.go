package main

func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] < arr[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

//func main() {
//	arr := []int{0, 1, 0}
//	fmt.Printf("%v \n", peakIndexInMountainArray(arr))
//}
