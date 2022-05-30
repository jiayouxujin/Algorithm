package main

import (
	"sort"
)

func twoSum(nums []int, begin, target int) [][]int {
	res := make([][]int, 0)
	left, right := begin, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		lo, hi := nums[left], nums[right]
		if sum < target {
			for left < right && nums[left] == lo {
				left++
			}
		} else if sum > target {
			for left < right && nums[right] == hi {
				right--
			}
		} else {
			tmp := []int{lo, hi}
			res = append(res, tmp)
			for left < right && nums[left] == lo {
				left++
			}
			for left < right && nums[right] == hi {
				right--
			}
		}
	}
	return res
}

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		tmp := twoSum(nums, i+1, 0-nums[i])
		for _, v := range tmp {
			v = append(v, nums[i])
			res = append(res, v)
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

//func main() {
//	nums := []int{-1, 0, 1, 2, -1, -4}
//	res:=threeSum(nums)
//	for i,_:=range res{
//		fmt.Printf("%v\n",res[i])
//	}
//}
