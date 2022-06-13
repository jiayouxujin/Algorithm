package main

func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1
	for left < right {
		mid := left + (right-left)/2
		if numbers[mid] < numbers[right] {
			right = mid
		} else if numbers[mid] > numbers[right] {
			left = mid + 1
		} else {
			right -= 1
		}
	}
	return numbers[left]
}

//func main() {
//	numbers := []int{3, 4, 5, 1, 2}
//	fmt.Printf("%v \n", minArray(numbers))
//}
