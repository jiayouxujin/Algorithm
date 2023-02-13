package main

func add(a int, b int) int {
	if b == 0 {
		return a
	}
	sum := a ^ b
	carry := (a & b) << 1
	return add(sum, carry)
}
