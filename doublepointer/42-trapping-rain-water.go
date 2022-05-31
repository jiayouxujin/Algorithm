package main

import (
	"fmt"
)

//version 1
//func trap(height []int) int {
//	n, res := len(height), 0
//	for i := 1; i < n-1; i++ {
//		l_max, r_max := 0, 0
//		for j := i; j < n; j++ {
//			r_max = max(r_max, height[j])
//		}
//		for j := i; j >= 0; j-- {
//			l_max = max(l_max, height[j])
//		}
//		res += min(r_max, l_max) - height[i]
//	}
//	return res
//}

//version 2
//func trap(height []int) int {
//	n, res := len(height), 0
//	l_max, r_max := make([]int, n), make([]int, n)
//	l_max[0], r_max[n-1] = height[0], height[n-1]
//
//	//从左往右
//	for i := 1; i < n; i++ {
//		l_max[i] = max(height[i], l_max[i-1])
//	}
//
//	//从右往左
//	for i := n - 2; i >= 0; i-- {
//		r_max[i] = max(height[i], r_max[i+1])
//	}
//
//	for i := 1; i < n-1; i++ {
//		res += min(l_max[i], r_max[i]) - height[i]
//	}
//	return res
//}

//version 3
func trap(height []int) int {
	left,right:=0,len(height)-1
	l_max,r_max,res:=0,0,0
	for left<right{
		l_max=max(height[left],l_max)
		r_max=max(height[right],r_max)

		if l_max<r_max{
			res+=l_max-height[left]
			left++
		}else{
			res+=r_max-height[right]
			right--
		}
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Printf("%v \n", trap(height))
}
