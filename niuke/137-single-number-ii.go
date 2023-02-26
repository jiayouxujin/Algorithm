package main

func singleNumber(nums []int) int {
	res := int32(0)
	for i := 0; i < 32; i++ {
		total := 0
		for _, num := range nums {
			total += ((num >> i) & 1)
		}
		if total%3 != 0 {
			res |= (1 << i)
		}
	}
	return int(res)
}
