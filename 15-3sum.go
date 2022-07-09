package main

import "sort"

func threeSum(nums []int) [][]int {
	if nums == nil || len(nums) < 3 {
		return [][]int{}
	}
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
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

func twoSum(nums []int, start, target int) [][]int {
	left, right, res := start, len(nums)-1, make([][]int, 0)
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
