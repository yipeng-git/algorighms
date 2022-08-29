package leetcode

func minPathSum(grid [][]int) int {
	path := make([][]int, len(grid))
	for i, _ := range path {
		path[i] = make([]int, len(grid[0]))
	}
	path[0][0] = grid[0][0]
	for i := 1; i < len(path[0]); i += 1 {
		path[0][i] = path[0][i-1] + grid[0][i]
	}
	for i := 1; i < len(path); i += 1 {
		path[i][0] = path[i-1][0] + grid[i][0]
	}
	for i := 1; i < len(path); i += 1 {
		for j := 1; j < len(path[0]); j += 1 {
			top := grid[i][j] + path[i-1][j]
			left := grid[i][j] + path[i][j-1]
			if top > left {
				path[i][j] = left
			} else {
				path[i][j] = top
			}
		}
	}
	return path[len(grid)-1][len(grid[0])-1]
}
