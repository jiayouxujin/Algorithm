package main

//接雨水 min(max(height[0..left]),max(height[left..right]))-height[i]
func trap(height []int) int {
	left, right := 0, len(height)-1
	lMax, rMax, res := 0, 0, 0
	for left < right {
		lMax = max(lMax, height[left])
		rMax = max(rMax, height[right])

		if lMax < rMax {
			res += lMax - height[left]
			left++
		} else {
			res += rMax - height[right]
			right--
		}
	}
	return res
}

//func main() {
//	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
//	fmt.Printf("%v \n", trap(height))
//}
