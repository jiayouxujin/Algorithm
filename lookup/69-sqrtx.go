package main

import (
	"math"
)

//version 1
//func mySqrt(x int) int {
//	left, right := 0, x
//	for left < right {
//		mid := (left + right + 1) >> 1
//		if mid*mid > x {
//			right = mid - 1
//		} else {
//			left = mid
//		}
//	}
//	return left
//}

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	C, x0 := float64(x), float64(x)
	for {
		xi := 0.5 * (x0 + C/x0)
		if math.Abs(xi-x0) < 1e-7 {
			break
		}
		x0 = xi
	}
	return int(x0)
}

//func main() {
//	fmt.Printf("%v \n", mySqrt(8))
//}
