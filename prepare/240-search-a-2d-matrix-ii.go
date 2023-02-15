package main

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := 0, len(matrix[0])-1
	for m < len(matrix) && n >= 0 {
		if matrix[m][n] == target {
			return true
		} else if matrix[m][n] > target {
			n--
		} else {
			m++
		}
	}
	return false
}
