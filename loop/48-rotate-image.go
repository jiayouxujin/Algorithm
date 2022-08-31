package main

func rotate(matrix [][]int) {
	if matrix == nil || len(matrix) == 0 {
		return
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if i > j {
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		left, right := 0, len(matrix[i])-1
		for left < right {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
			left++
			right--
		}
	}
}
