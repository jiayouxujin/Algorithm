/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */

// @lc code=start
var dir [4][2]int=[4][2]int{{0,1},{0,-1},{1,0},{-1,0}}

func numIslands(grid [][]byte) int {
	var res int
	for i:=range grid{
		for j:=range grid[i]{
			if grid[i][j]=='1'{
				dfs(grid,i,j)
				res++
			}
		}
	}
	return res
}

func dfs(grid [][]byte,i,j int){
	if i<0||i>len(grid)-1||j<0||j>len(grid[0])-1{
		return 
	}

	if grid[i][j]=='1'{
		grid[i][j]='0'
		for k:=range dir{
			dfs(grid,i+dir[k][0],j+dir[k][1])
		}
	}
}
// @lc code=end

