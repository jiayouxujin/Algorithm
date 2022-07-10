package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		tmp := target - nums[i]
		if v, ok := m[tmp]; ok {
			return []int{v, i}
		}
		m[nums[i]] = i
	}
	return []int{}
}

//func main() {
//	nums, k := []int{2, 7, 11, 15}, 9
//	fmt.Printf("%v \n", twoSum(nums, k))
//}
