package main

import "fmt"

func isValid(s string) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else if (s[i] == ')' && len(stack) > 0 && stack[len(stack)-1] == '(') ||
			(s[i] == '}' && len(stack) > 0 && stack[len(stack)-1] == '{') ||
			(s[i] == ']' && len(stack) > 0 && stack[len(stack)-1] == '[') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

func main() {
	s := "()"
	fmt.Printf("%v \n", isValid(s))
}
