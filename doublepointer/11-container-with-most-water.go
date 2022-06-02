package main

func maxArea(height []int) int {
	res, left, right := 0, 0, len(height)-1
	for left < right {
		cur_area:=min(height[right],height[left])*(right-left)
		res=max(res,cur_area)
		if height[right]<height[left] {
			right--
		}else{
			left++
		}
	}
	return res
}


//func main() {
//	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
//	fmt.Printf("%v \n", maxArea(height))
//}
