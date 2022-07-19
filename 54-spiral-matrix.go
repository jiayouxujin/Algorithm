package main

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	upper, lower, left, right := 0, m-1, 0, n-1
	res := make([]int, 0)
	for len(res) < m*n {
		if upper <= lower {
			//从左往右
			for i := left; i <= right; i++ {
				res = append(res, matrix[upper][i])
			}
			upper++
		}

		if left <= right {
			//从上往下
			for j := upper; j <= lower; j++ {
				res = append(res, matrix[j][right])
			}
			right--
		}

		if upper <= lower {
			//从右往左
			for i := right; i >= left; i-- {
				res = append(res, matrix[lower][i])
			}
			lower--
		}

		if left <= right {
			//从下往上
			for j := lower; j >= upper; j-- {
				res = append(res, matrix[j][left])
			}
			left++
		}
	}
	return res
}

//func main() {
//	matrix := [][]int{
//		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
//	}
//	fmt.Printf("%v \n", spiralOrder(matrix))
//}
