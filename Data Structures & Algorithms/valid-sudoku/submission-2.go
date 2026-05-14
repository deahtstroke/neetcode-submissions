func isValidSudoku(board [][]byte) bool {
 // Naive approach:
 // Go first through each row to check if there's duplicates or not
	for i := 0; i < 9; i++ {
		hashset := make(map[byte]bool, 9)
		for j := 0; j < 9; j++ {
			curr := board[i][j]
			if curr == '.' {
				continue
			}
			if  hashset[curr] {
				return false
			}
			hashset[curr] = true
		}
	}
	
	// Go now through each column
	for i := 0; i < 9; i++ {
		hashset := make(map[byte]bool, 9)
		for j := 0; j < 9; j++ {
			curr := board[j][i]
			if curr == '.' {
				continue
			}
			if hashset[curr] {
				return false
			}
			hashset[curr] = true
		}		
	}

	// Go now through each 3x3 grid on the 9x9 grid
	for i := 0; i < 9; i++ {
		//  origin has to reflect the (0,0) x and y position of the current square
		// (row / 3) * 3 + (col / 3)
		hashset := make(map[byte]bool, 9)
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				row := (i / 3) * 3 + j
				col := (i % 3) * 3 + k
				curr := board[row][col]
				if curr == '.' {
					continue
				}
				if hashset[curr] {
					return false
				}

				hashset[curr] = true
			}
		}
	}

	return true
}
