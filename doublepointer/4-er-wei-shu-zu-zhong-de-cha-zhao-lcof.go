package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	col, row := len(matrix[0])-1, 0
	for col >= 0 && row < len(matrix) {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}
	return false
}
//func main() {
//	matrix := [][]int{
//		{1, 4, 7, 11, 15},
//		{2, 5, 8, 12, 19},
//		{3, 6, 9, 16, 22},
//		{10, 13, 14, 17, 24},
//		{18, 21, 23, 26, 30},
//	}
//	fmt.Printf("%v \n", findNumberIn2DArray(matrix, 20))
//}
