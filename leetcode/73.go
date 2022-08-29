package leetcode

func setZeroes(matrix [][]int) {
	rows := map[int]struct{}{}
	cols := map[int]struct{}{}
	for i := 0; i < len(matrix); i += 1 {
		for j := 0; j < len(matrix[0]); j += 1 {
			if matrix[i][j] == 0 {
				rows[i] = struct{}{}
				cols[j] = struct{}{}
			}
		}
	}
	for row, _ := range rows {
		for i, _ := range matrix[row] {
			matrix[row][i] = 0
		}
	}
	for col, _ := range cols {
		for i, _ := range matrix {
			matrix[i][col] = 0
		}
	}
}
