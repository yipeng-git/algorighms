package leetcode

func uniquePaths(m int, n int) int {
	path := make([][]int, m)
	for i, _ := range path {
		path[i] = make([]int, n)
	}
	for i, _ := range path[0] {
		path[0][i] = 1
	}
	for i, _ := range path {
		path[i][0] = 1
	}
	for i := 1; i < len(path); i += 1 {
		for j := 1; j < len(path[0]); j += 1 {
			path[i][j] = path[i-1][j] + path[i][j-1]
		}
	}
	return path[m-1][n-1]
}
