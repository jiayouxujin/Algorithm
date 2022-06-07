package main

func moveZeroes(nums []int) {
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != 0 {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	for left < len(nums) {
		nums[left] = 0
		left++
	}
}

//func main() {
//	nums := []int{0, 1, 0, 3, 12}
//	moveZeroes(nums)
//}
