package main

func isValid(s string) bool {
	stack := make([]byte, 0)
	for _, b := range s {
		if b == '{' || b == '[' || b == '(' {
			stack = append(stack, byte(b))
		} else if len(stack) > 0 {
			if b == '}' && stack[len(stack)-1] != '{' {
				return false
			} else if b == ']' && stack[len(stack)-1] != '[' {
				return false
			} else if b == ')' && stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
