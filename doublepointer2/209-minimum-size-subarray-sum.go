package main

func minSubArrayLen(target int, nums []int) int {
	left, right, res, sum := 0, 0, len(nums)+1, 0
	for right < len(nums) {
		sum += nums[right]
		right++

		for sum >= target {
			if right-left < res {
				res = right - left
			}
			sum -= nums[left]
			left++
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

//func main() {
//	target, nums := 7, []int{2, 3, 1, 2, 4, 3}
//	fmt.Printf("%v \n", minSubArrayLen(target, nums))
//}
