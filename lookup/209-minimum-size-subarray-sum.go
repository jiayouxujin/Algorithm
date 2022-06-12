package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	left, res, sum := 0, len(nums)+1, 0
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target {
			if right-left+1 < res {
				res = right - left + 1
			}
			sum -= nums[left]
			left++
		}
	}
	if res != len(nums)+1 {
		return res
	}
	return 0
}

func main() {
	target, nums := 7, []int{2, 3, 1, 2, 4, 3}
	fmt.Printf("%v \n", minSubArrayLen(target, nums))
}
