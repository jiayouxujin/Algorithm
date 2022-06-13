package main

import "fmt"

func searchMatrix2(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	col, row := len(matrix[0])-1, 0
	for col >= 0 && row < len(matrix) {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			row++
		} else {
			col--
		}
	}
	return false
}

func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target := 3
	fmt.Printf("%v \n", searchMatrix2(matrix, target))
}
