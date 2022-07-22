package main

// 1 2 3 4
// 5 6 7 8
// 9 10 11 12
// 13 14 15 16
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	res := make([]int, 0)
	upper, lower, left, right := 0, m-1, 0, n-1
	for len(res) < m*n {
		if upper <= lower {
			for i := left; i <= right; i++ {
				res = append(res, matrix[upper][i])
			}
			upper++
		}
		if left <= right {
			for i := upper; i <= lower; i++ {
				res = append(res, matrix[i][right])
			}
			right--
		}
		if upper <= lower {
			for i := right; i >= left; i-- {
				res = append(res, matrix[lower][i])
			}
			lower--
		}
		if left <= right {
			for i := lower; i >= upper; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}
	return res
}

//func main() {
//	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
//	fmt.Printf("%v \n", spiralOrder(matrix))
//}
