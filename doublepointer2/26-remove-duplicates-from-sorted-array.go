package main

func removeDuplicates(nums []int) int {
	left, right := 0, 1
	for right < len(nums) {
		if nums[right] != nums[left] {
			left++
			nums[left] = nums[right]
		}
		right++
	}
	return left + 1
}

//func main() {
//	nums := []int{1, 1, 2}
//	fmt.Printf("%v \n", removeDuplicates(nums))
//}
