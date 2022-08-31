package main

func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	up, down, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1

	for len(res) < len(matrix)*len(matrix[0]) {
		//上边界
		if up <= down {
			for i := left; i <= right; i++ {
				res = append(res, matrix[up][i])
			}
			up++
		}
		//右边界
		if right >= left {
			for i := up; i <= down; i++ {
				res = append(res, matrix[i][right])
			}
			right--
		}
		//下边界
		if up <= down {
			for i := right; i >= left; i-- {
				res = append(res, matrix[down][i])
			}
			down--
		}
		//左边界
		if left <= right {
			for i := down; i >= up; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}
	return res
}
