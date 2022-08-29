package leetcode

func spiralOrder(matrix [][]int) []int {
	i := 0 // row
	j := 0 // col
	visited := make([][]bool, len(matrix))
	for k, _ := range visited {
		visited[k] = make([]bool, len(matrix[0]))
	}
	res := make([]int, len(matrix)*len(matrix[0]))
	dir := 0
	for n := 0; n < len(matrix)*len(matrix[0]); n += 1 {
		res[n] = matrix[i][j]
		visited[i][j] = true
		if dir == 0 {
			// go right
			if j == len(matrix[0])-1 || visited[i][j+1] {
				dir = 1
				i += 1
			} else {
				j += 1
			}
		} else if dir == 1 {
			// go down
			if i == len(matrix)-1 || visited[i+1][j] {
				dir = 2
				j -= 1
			} else {
				i += 1
			}
		} else if dir == 2 {
			// go left
			if j == 0 || visited[i][j-1] {
				dir = 3
				i -= 1
			} else {
				j -= 1
			}
		} else if dir == 3 {
			// go up
			if i == 0 || visited[i-1][j] {
				dir = 0
				j += 1
			} else {
				i -= 1
			}
		}
	}
	return res
}
