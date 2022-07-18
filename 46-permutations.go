package main

// 回溯
func permute(nums []int) [][]int {
	res, tmp, used := make([][]int, 0), make([]int, 0), make([]bool, len(nums))
	backtrackPermute(nums, tmp, &used, &res)
	return res
}

func backtrackPermute(nums, tmp []int, used *[]bool, res *[][]int) {
	if len(tmp) == len(nums) {
		forRes := make([]int, len(tmp))
		copy(forRes, tmp)
		*res = append(*res, forRes)
		return
	}
	for i := 0; i < len(nums); i++ {
		if (*used)[i] {
			continue
		}
		(*used)[i] = true
		tmp = append(tmp, nums[i])
		backtrackPermute(nums, tmp, used, res)
		tmp = tmp[:len(tmp)-1]
		(*used)[i] = false
	}
}

//func main() {
//	nums := []int{1, 2, 3}
//	fmt.Printf("%v \n", permute(nums))
//}
