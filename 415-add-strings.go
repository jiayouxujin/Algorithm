package main

import "strconv"

func addStrings(num1 string, num2 string) string {
	res := ""
	i, j, carry := len(num1)-1, len(num2)-1, 0
	var n1, n2 int
	for i >= 0 || j >= 0 {
		n1, n2 = 0, 0
		if i >= 0 {
			n1 = int(num1[i] - '0')
		}
		if j >= 0 {
			n2 = int(num2[j] - '0')
		}
		sum := n1 + n2 + carry
		carry = sum / 10
		res = strconv.Itoa(sum%10) + res
		i--
		j--
	}
	if carry > 0 {
		return "1" + res
	}
	return res
}
