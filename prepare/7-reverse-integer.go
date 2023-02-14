package main

import "strconv"

func reverse2(x int) int {
	maxValue, minValue := 1<<31-1, -1<<31
	res, str := 0, strconv.Itoa(x)
	if str[0] == '-' {
		res, _ = strconv.Atoi("-" + reverseStr(str[1:]))
	} else {
		res, _ = strconv.Atoi(reverseStr(str))
	}
	if res < minValue || res > maxValue {
		res = 0
	}
	return res
}

func reverseStr(str string) string {
	arr := []rune(str)
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
	return string(arr)
}
