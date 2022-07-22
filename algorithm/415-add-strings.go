package main

import (
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	if num1 == "" {
		return num2
	}
	if num2 == "" {
		return num1
	}

	i, j, carray := len(num1)-1, len(num2)-1, 0
	res := ""
	for i >= 0 || j >= 0 {
		m, n := 0, 0
		if i >= 0 {
			m = int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			n = int(num2[j] - '0')
			j--
		}
		sum := m + n + carray
		carray = sum / 10
		res = strconv.Itoa(sum%10) + res
	}
	if carray > 0 {
		return "1" + res
	}
	return res
}

//func main() {
//	num1, num2 := "11", "123"
//	fmt.Printf("%v \n", addStrings(num1, num2))
//}
