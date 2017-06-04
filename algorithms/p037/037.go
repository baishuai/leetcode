package p037

/**
Write a program to solve a Sudoku puzzle by filling the empty cells.

Empty cells are indicated by the character '.'.

You may assume that there will be only one unique solution.
*/

type bit9Set int16

func (bs *bit9Set) contains(b byte) bool {
	return (*bs & (1 << b)) > 0
}

func (bs *bit9Set) insert(b byte) {
	*bs = *bs | (1 << b)
}

func (bs *bit9Set) remove(b byte) {
	*bs = *bs &^ (1 << b)
}

func solveSudoku(board [][]byte) {
	rows := make([]bit9Set, 9)
	cols := make([]bit9Set, 9)
	grids := make([]bit9Set, 9)
	dotCount := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			b := board[i][j]
			if b == '.' {
				dotCount++
			} else {
				b = b - '0'
				rows[i].insert(b)
				cols[j].insert(b)
				grids[i/3*3+j/3].insert(b)
			}
		}
	}

	var process func(r byte) bool
	process = func(r byte) bool {

		var i, j byte
	out:
		for i = r; i < 9; i++ {
			for j = 0; j < 9; j++ {
				if board[i][j] == '.' {
					break out
				}
			}
		}

		if i == 9 {
			return true
		}
		find := false
		for v := byte(1); v <= 9; v++ {
			if !rows[i].contains(v) &&
				!cols[j].contains(v) &&
				!grids[i/3*3+j/3].contains(v) {
				rows[i].insert(v)
				cols[j].insert(v)
				grids[i/3*3+j/3].insert(v)
				board[i][j] = '0' + v
				if process(i) {
					find = true
					break
				}

				rows[i].remove(v)
				cols[j].remove(v)
				grids[i/3*3+j/3].remove(v)
			}
		}

		if find {
			return true
		} else {
			board[i][j] = '.'
			return false
		}
	}

	process(0)
}
