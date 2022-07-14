package main

func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res += 1
				dfs(grid, i, j)
			}
		}
	}
	return res
}

func dfs(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return
	}
	if grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}

//func main() {
//	grid := [][]byte{
//		{'1', '1', '1', '1', '0'},
//		{'1', '1', '0', '1', '0'},
//		{'1', '1', '0', '0', '0'},
//		{'0', '0', '0', '0', '0'},
//	}
//	fmt.Printf("%v \n", numIslands(grid))
//}
