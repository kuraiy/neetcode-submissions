func isValidSudoku(board [][]byte) bool {
	for i := range board {
		if !checkRow(board[i]) {
			return false
		}
	}

	for i := range 9 {
		if !checkColumn(board, i) {
			return false
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			if !checkBox(board, i, j) {
				return false
			}
		}
	}

	return true
}

func checkRow(row []byte) bool {
	dots := 0
	nums := make(map[byte]struct{})

	for _, v := range row {
		if v == '.' {
			dots++
		} else {
			if _, ok := nums[v]; !ok {
				nums[v] = struct{}{}
			}
		}
	}

	return dots+len(nums) == 9
}

func checkColumn(board [][]byte, ind int) bool {
	dots := 0
	nums := make(map[byte]struct{})

	for i := range 9 {
		if board[i][ind] == '.' {
			dots++
		} else {
			if _, ok := nums[board[i][ind]]; !ok {
				nums[board[i][ind]] = struct{}{}
			}
		}
	}

	return dots+len(nums) == 9
}

func checkBox(box [][]byte, rowInd, columdInd int) bool {
	dots := 0
	nums := make(map[byte]struct{})

	for i := rowInd; i < rowInd+3; i++ {
		for j := columdInd; j < columdInd+3; j++ {
			if box[i][j] == '.' {
				dots++
			} else {
				if _, ok := nums[box[i][j]]; !ok {
					nums[box[i][j]] = struct{}{}
				}
			}
		}
	}

	return dots+len(nums) == 9
}
