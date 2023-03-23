package valid_sudoku

func isValidSudoku(board [][]byte) bool {
	var rows, columns [9][9]int
	var subboxes [3][3][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			v := board[i][j]
			if v == '.' {
				continue
			}
			index := v - '1'
			rows[i][index]++
			columns[j][index]++
			subboxes[i/3][j/3][index]++
			// 正常情况下 每行、每列、每个九宫格内的每个数字均只有1个，如果大于1则代表不正确
			if rows[i][index] > 1 || columns[j][index] > 1 || subboxes[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true
}
