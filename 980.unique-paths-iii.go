/*
 * @lc app=leetcode id=980 lang=golang
 *
 * [980] Unique Paths III
 */

// @lc code=start
var dir=[][]int{
	{1,0},
	{0,1},
	{-1,0},
	{0,-1},
}

func uniquePathsIII(grid [][]int) int {
	visited:=make([][]bool,len(grid))
	for i:=0;i<len(visited);i++{
		visited[i]=make([]bool,len(grid[0]))
	}

	empty,startx,starty,endx,endy,res,path:=0,0,0,0,0,0,[]int{}
	for i,v:=range grid{
		for j,vv:=range v{
			switch vv {
			case 0:
				empty++
			case 1:
				startx,starty=i,j
			case 2:
				endx,endy=i,j
			}
		}
	}

	findUniquePath(grid,visited,empty+1,startx,starty,endx,endy,path,&res)
	return res
}

func isInPath(grid [][]int,x,y int)bool{
	return x>=0&&x<len(grid)&&y>=0&&y<len(grid[0])
}

func findUniquePath(board [][]int,visited [][]bool,empty,startx,starty,endx,endy int,path []int,res *int){
	if startx==endx&&starty==endy{
		if empty==0{
			*res++
		}
		return 
	}

	if board[startx][starty]>=0{
		visited[startx][starty]=true
		empty--
		path=append(path,startx)
		path=append(path,starty)
		for i:=0;i<len(dir);i++{
			nx:=startx+dir[i][0]
			ny:=starty+dir[i][1]
			if isInPath(board,nx,ny)&&!visited[nx][ny]{
				findUniquePath(board,visited,empty,nx,ny,endx,endy,path,res)
			}
		}
		visited[startx][starty]=false
		path=path[:len(path)-2]
	}
	return 
}
// @lc code=end

