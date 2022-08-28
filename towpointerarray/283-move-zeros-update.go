package main

import "fmt"

func moveZero2(nums []int) {
	curIndex := 0
	for i, v := range nums {
		if v != 1 && v != 3 && v != 6 {
			curIndex = i
		}
		for curIndex > 0 && (nums[curIndex-1] == 1 || nums[curIndex-1] == 3 || nums[curIndex-1] == 6) {
			nums[curIndex], nums[curIndex-1] = nums[curIndex-1], nums[curIndex]
			curIndex -= 1
		}
	}
	fmt.Printf("%v \n", nums)
}

//func main() {
//	nums := []int{0, 1, 6, 3, 6, 3, 12}
//	moveZero2(nums)
//}
