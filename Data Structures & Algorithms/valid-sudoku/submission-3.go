func isValidSudoku(board [][]byte) bool {
	type Pair struct {
		X int
		Y int
	}

	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	sqrs := make(map[Pair]map[byte]bool, 9) // key is Pair{X, Y}

	// Initialize hash sets for each of the 9 rows and columns
	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool, 9)
		cols[i] = make(map[byte]bool, 9)
	}
	for i, row := range board {
		for j, val := range row {
			if val == '.' {
				continue
			}

			box := Pair{X: i / 3, Y: j / 3}
			
			if sqrs[box] == nil {
				sqrs[box] = make(map[byte]bool, 9)
			}

			if cols[j][val] || rows[i][val] || sqrs[box][val] {
				return false
			}

			cols[j][val] = true
			rows[i][val] = true
			sqrs[box][val] = true
		}
	}	

	return true
}
