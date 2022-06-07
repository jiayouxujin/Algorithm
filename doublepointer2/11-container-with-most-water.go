package main

func maxArea(height []int) int {
	left, right, res := 0, len(height)-1, 0
	for left < right {
		curArea := min(height[left], height[right]) * (right - left)
		if curArea > res {
			res = curArea
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
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

//func main() {
//	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
//	fmt.Printf("%v \n", maxArea(height))
//}
