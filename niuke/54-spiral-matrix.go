package main

func spiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	res, left, right, up, down := make([]int, 0), 0, len(matrix[0])-1, 0, len(matrix)-1
	for len(res) < len(matrix)*len(matrix[0]) {
		if up <= down {
			for i := left; i <= right; i++ {
				res = append(res, matrix[up][i])
			}
			up++
		}
		if right >= left {
			for j := up; j <= down; j++ {
				res = append(res, matrix[j][right])
			}
			right--
		}
		if up <= down {
			for i := right; i >= left; i-- {
				res = append(res, matrix[down][i])
			}
			down--
		}
		if left <= right {
			for j := down; j >= up; j-- {
				res = append(res, matrix[j][left])
			}
			left++
		}
	}
	return res
}
