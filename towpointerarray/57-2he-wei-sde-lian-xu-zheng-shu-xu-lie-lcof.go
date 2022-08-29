package main

import "fmt"

func findContinuousSequence(target int) [][]int {
	if target <= 1 {
		return [][]int{}
	}

	left, right, tmp, ans := 1, 1, 1, make([][]int, 0)
	for right < target {
		right++
		tmp += right

		for tmp > target {
			tmp -= left
			left++
		}

		if tmp == target && left != right {
			res := []int{}
			for i := left; i <= right; i++ {
				res = append(res, i)
			}
			ans = append(ans, res)
		}

	}
	return ans
}

func main() {
	fmt.Printf("%v \n", findContinuousSequence(9))
}
