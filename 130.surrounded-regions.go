/*
 * @lc app=leetcode id=130 lang=golang
 *
 * [130] Surrounded Regions
 */

// @lc code=start
var dir [4][2]int=[4][2]int{{0,1},{0,-1},{1,0},{-1,0}}

func solve(board [][]byte)  {
    for i:=range board{
		for j:=range board[i]{
			if i==0||i==len(board)-1||j==0||j==len(board[i])-1{
				if board[i][j]=='O'{
					dfs(i,j,board)
				}
			}
		}
	}

	for i:=range board{
		for j:=range board[i]{
			if board[i][j]=='*'{
				board[i][j]='O'
			}else if board[i][j]=='O'{
				board[i][j]='X'
			}
		}
	}
}

func dfs(i,j int,board [][]byte){
	if i<0||i>len(board)-1||j<0||j>len(board[0])-1{
		return 
	}

	if board[i][j]=='O'{
		board[i][j]='*'
		for k:=range dir{
			dfs(i+dir[k][0],j+dir[k][1],board)
		}
	}

}
// @lc code=end

