package main

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	return solve(nums, left, right, target)
}

func solve(nums []int, l, r, target int) int {
	if l >= r {
		if nums[l] == target {
			return l
		}
		return -1
	}

	mid := l + (r-l)/2
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return solve(nums, l, mid-1, target)
	} else {
		return solve(nums, mid, r, target)
	}
}

//func main() {
//	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
//	fmt.Printf("%v \n", search(nums, 8))
//}
