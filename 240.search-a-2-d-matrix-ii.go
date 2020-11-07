/*
 * @lc app=leetcode id=240 lang=golang
 *
 * [240] Search a 2D Matrix II
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix)==0{
		return false
	}

	row,col:=0,len(matrix[0])-1
	for row<=len(matrix)-1&&col>=0{
		if target==matrix[row][col]{
			return true
		}else if target>matrix[row][col]{
			row++
		}else{
			col--
		}
	}
	return false
}
// @lc code=end

