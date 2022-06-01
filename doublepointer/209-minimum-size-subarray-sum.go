package main

func minSubArrayLen(target int, nums []int) int {
	left, right, res, cur := 0, 0, len(nums)+1, 0
	for right < len(nums) {
		cur += nums[right]
		right++

		for cur >= target {
			if right-left< res {
				res = right - left
			}
			cur -= nums[left]
			left++
		}
	}
	if res==len(nums)+1{
		return 0
	}
	return res
}

//func main() {
//	target, nums := 7, []int{2, 3, 1, 2, 4, 3}
//	fmt.Println(minSubArrayLen(target, nums))
//}
