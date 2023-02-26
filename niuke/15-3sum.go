package main

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		tows := towSum(nums, i+1, -nums[i])
		for _, t := range tows {
			tmp := []int{nums[i]}
			tmp = append(tmp, t...)
			res = append(res, tmp)
		}
		for i < len(nums)-2 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

func towSum(nums []int, start, target int) [][]int {
	res := make([][]int, 0)
	left, right := start, len(nums)-1
	for left < right {
		lo, hi := nums[left], nums[right]
		sum := lo + hi
		if sum == target {
			tmp := []int{lo, hi}
			res = append(res, tmp)

			for left < right && nums[left] == lo {
				left++
			}

			for left < right && nums[right] == hi {
				right--
			}
		} else if sum > target {
			for right > left && nums[right] == hi {
				right--
			}
		} else {
			for left < right && nums[left] == lo {
				left++
			}
		}
	}
	return res
}
