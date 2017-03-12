package p036

/**
Determine if a Sudoku is valid, according to: Sudoku Puzzles - The Rules.

The Sudoku board could be partially filled,
 where empty cells are filled with the character '.'.

A valid Sudoku board (partially filled) is not necessarily solvable.
 Only the filled cells need to be validated.
 */

type bit9Set int16

func (bs *bit9Set) containsOrInsert(b byte) bool {
	contains := (*bs & (1 << b)) > 0
	*bs = *bs | (1 << b)
	return contains
}

func isValidSudoku(board [][]byte) bool {
	row := bit9Set(0)
	cols := make([]bit9Set, 9)
	grids := make([]bit9Set, 3)
	ans := true
	for i := 0; i < 9; i++ {
		row = 0
		if i%3 == 0 {
			grids[0], grids[1], grids[2] = 0, 0, 0
		}
		for j := 0; j < 9; j++ {
			b := board[i][j]
			if b == '.' {
				continue
			}
			b = b - '0'
			if row.containsOrInsert(b) ||
				cols[j].containsOrInsert(b) ||
				grids[j/3].containsOrInsert(b) {
				ans = false
				break
			}
		}
	}
	return ans
}
