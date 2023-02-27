package main

func trap(height []int) int {
	left, right, res, lMax, rMax := 0, len(height)-1, 0, 0, 0
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
