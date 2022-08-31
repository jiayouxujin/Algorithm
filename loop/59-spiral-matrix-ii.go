package main

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	up, down, left, right, num := 0, n-1, 0, n-1, 1
	for num <= n*n {
		if up <= down {
			for i := left; i <= right; i++ {
				matrix[up][i] = num
				num++
			}
			up++
		}
		if right >= left {
			for i := up; i <= down; i++ {
				matrix[i][right] = num
				num++
			}
			right--
		}
		if up <= down {
			for i := right; i >= left; i-- {
				matrix[down][i] = num
				num++
			}
			down--
		}
		if left <= right {
			for i := down; i >= up; i-- {
				matrix[i][left] = num
				num++
			}
			left++
		}
	}
	return matrix
}
