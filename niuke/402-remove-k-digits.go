package main

import "strconv"

func removeKdigits(num string, k int) string {
	if k == len(num) {
		return "0"
	}
	stack := make([]int, 0)
	for i := 0; i < len(num); i++ {
		c := int(num[i] - '0')
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > c {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, c)
	}
	stack = stack[:len(stack)-k]
	res := ""
	for len(stack) > 1 && stack[0] == 0 {
		stack = stack[1:]
	}
	for i := 0; i < len(stack); i++ {
		res += strconv.Itoa(stack[i])
	}
	return res
}
