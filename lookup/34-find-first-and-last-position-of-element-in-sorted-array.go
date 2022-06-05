package main

func searchRange(nums []int, target int) []int {
	return []int{left(nums, target), right(nums, target)}
}

func left(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			//收缩右侧
			r = mid - 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if l >= len(nums) || nums[l] != target {
		return -1
	}
	return l
}

func right(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			//收缩左侧
			l = mid + 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if r < 0 || nums[r] != target {
		return -1
	}
	return r
}

//func main() {
//	nums, target := []int{5, 7, 7, 8, 8, 10}, 8
//	_ = searchRange(nums, target)
//}
