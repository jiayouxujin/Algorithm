package main

func spiralOrder(matrix [][]int) []int {
	var res []int
	up, down, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	for len(res) < len(matrix)*len(matrix[0]) {
		// 从左往右
		if up <= down {
			for i := left; i <= right; i++ {
				res = append(res, matrix[up][i])
			}
			up++
		}
		// 从上到下
		if left <= right {
			for i := up; i <= down; i++ {
				res = append(res, matrix[i][right])
			}
			right--
		}
		//从右到左
		if up <= down {
			for i := right; i >= left; i-- {
				res = append(res, matrix[down][i])
			}
			down--
		}
		//从下到上
		if left <= right {
			for i := down; i >= up; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}
	return res
}
