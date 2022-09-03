package main

func splitArray(nums []int, m int) int {
	lo, hi := getMax(nums), getSum(nums)
	for lo < hi {
		mid := lo + (hi-lo)/2
		n := getSplit(nums, mid)
		if n == m {
			hi = mid
		} else if n < m {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func getSplit(nums []int, max int) int {
	count, sum := 1, 0
	for i := 0; i < len(nums); i++ {
		if sum+nums[i] > max {
			count++
			sum = nums[i]
		} else {
			sum += nums[i]
		}
	}
	return count
}
func getMax(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if res < nums[i] {
			res = nums[i]
		}
	}
	return res
}

func getSum(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}
