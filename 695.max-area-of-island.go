/*
 * @lc app=leetcode id=695 lang=golang
 *
 * [695] Max Area of Island
 */

// @lc code=start
var dir [4][2]int=[4][2]int{{0,1},{0,-1},{1,0},{-1,0}}

func maxAreaOfIsland(grid [][]int) int {
	var res int
	for i:=range grid{
		for j:=range grid[i]{
			if grid[i][j]==1{
				area:=dfs(grid,i,j)
				if res<area{
					res=area
				}
			}
		}
	}	
	return res
}

func dfs(grid [][]int,i,j int)int{
	if i<0||i>len(grid)-1||j<0||j>len(grid[0])-1||grid[i][j]==0{
		return 0
	}
	grid[i][j]=0
	num:=1
	for k:=range dir{
		num+=dfs(grid,i+dir[k][0],j+dir[k][1])
	}
	return num
}
// @lc code=end

