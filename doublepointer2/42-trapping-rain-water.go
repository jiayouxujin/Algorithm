package main

func trap(height []int) int {
	left, right := 0, len(height)-1
	l_max, r_max, res := 0, 0, 0
	for left < right {
		l_max = max(height[left], l_max)
		r_max = max(height[right], r_max)

		if r_max < l_max {
			res += r_max - height[right]
			right--
		} else {
			res += l_max - height[left]
			left++
		}
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

//func main() {
//	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
//	fmt.Printf("%v \n", trap(height))
//}
