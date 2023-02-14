package main

func decodeString(s string) string {
	stack, numStack := make([]string, 0), make([]int, 0)
	res, num := "", 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
		} else if s[i] == '[' {
			numStack = append(numStack, num)
			stack = append(stack, res)
			res = ""
			num = 0
		} else if s[i] == ']' {
			temp := stack[len(stack)-1]
			count := numStack[len(numStack)-1]
			stack = stack[:len(stack)-1]
			numStack = numStack[:len(numStack)-1]
			for count > 0 {
				temp += res
				count--
			}
			res = temp
		} else {
			res += string(s[i])
		}
	}
	return res
}
