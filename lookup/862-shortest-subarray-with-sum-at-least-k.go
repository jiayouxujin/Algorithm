package main

func shortestSubarray(nums []int, k int) int {
	res, prefixSum := len(nums)+1, make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}

	deque := []int{}
	for i := range prefixSum {
		for len(deque) > 0 && prefixSum[i]-prefixSum[deque[0]] >= k {
			length := i - deque[0]
			if res > length {
				res = length
			}
			deque = deque[1:]
		}
		for len(deque) > 0 && prefixSum[i] <= prefixSum[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
	}
	if res <= len(nums) {
		return res
	}
	return -1
}

//func main() {
//	nums, k := []int{2, -1, 2}, 3
//	fmt.Printf("%v \n", shortestSubarray(nums, k))
//}
