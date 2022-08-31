package main

import "sort"

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 3 {
		return res
	}

	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		tow := towSum(nums, i+1, -nums[i])
		for _, v := range tow {
			tmp := append(v, nums[i])
			res = append(res, tmp)
		}

		for i < len(nums)-1 && nums[i] == nums[i+1] {
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
		if target == sum {
			res = append(res, []int{lo, hi})

			for left < right && nums[left] == lo {
				left++
			}
			for left < right && nums[right] == hi {
				right--
			}
		} else if target > sum {
			for left < right && nums[left] == lo {
				left++
			}
		} else {
			for left < right && nums[right] == hi {
				right--
			}
		}
	}
	return res
}
