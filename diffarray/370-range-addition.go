package main

import "fmt"

func getModifiedArray(length int, updates [][]int) []int {
	diff := make([]int, length)
	for i := 0; i < len(updates); i++ {
		diff[updates[i][0]] += updates[i][2]
		if updates[i][1]+1 < length {
			diff[updates[i][1]+1] -= updates[i][2]
		}
	}
	for i := 1; i < length; i++ {
		diff[i] += diff[i-1]
	}
	return diff
}

func main() {
	length, updates := 5, [][]int{
		{1, 3, 2},
		{2, 4, 3},
		{0, 2, -2},
	}
	fmt.Printf("%v \n", getModifiedArray(length, updates))
}
