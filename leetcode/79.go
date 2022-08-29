package leetcode

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i += 1 {
		for j := 0; j < len(board[0]); j += 1 {
			if existHelper(board, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func existHelper(board [][]byte, word string, idx, i, j int) bool {
	if idx >= len(word) {
		return true
	}
	if i >= len(board) || i < 0 || j < 0 || j >= len(board[0]) || board[i][j] == byte(0) || board[i][j] != word[idx] {
		return false
	}
	curr := board[i][j]
	board[i][j] = byte(0)
	a := existHelper(board, word, idx+1, i+1, j)
	if a == true {
		board[i][j] = curr
		return true
	}
	b := existHelper(board, word, idx+1, i-1, j)
	if b == true {
		board[i][j] = curr
		return true
	}
	c := existHelper(board, word, idx+1, i, j-1)
	if c == true {
		board[i][j] = curr
		return true
	}
	d := existHelper(board, word, idx+1, i, j+1)
	if d == true {
		board[i][j] = curr
		return true
	}
	board[i][j] = curr
	return false
}
