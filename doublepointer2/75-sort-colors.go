package main

func sortColors(nums []int) {
	zero, one := 0, 0
	for i, n := range nums {
		nums[i] = 2
		if n <= 1 {
			nums[one] = 1
			one++
		}
		if n == 0 {
			nums[zero] = 0
			zero++
		}
	}
}

//func main() {
//	nums := []int{2, 0, 2, 1, 1, 0}
//	sortColors(nums)
//}
