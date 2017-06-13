package p289

/**
According to the Wikipedia's article: "The Game of Life, also known simply as Life, is a cellular automaton devised by the British mathematician John Horton Conway in 1970."

Given a board with m by n cells, each cell has an initial state live (1) or dead (0). Each cell interacts with its eight neighbors (horizontal, vertical, diagonal) using the following four rules (taken from the above Wikipedia article):

Any live cell with fewer than two live neighbors dies, as if caused by under-population.
Any live cell with two or three live neighbors lives on to the next generation.
Any live cell with more than three live neighbors dies, as if by over-population..
Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
Write a function to compute the next state (after one update) of the board given its current state.

Follow up:
Could you solve it in-place? Remember that the board needs to be updated at the same time: You cannot update some cells first and then use their updated values to update other cells.
In this question, we represent the board using a 2D array. In principle, the board is infinite, which would cause problems when the active area encroaches the border of the array. How would you address these problems?

*/

/**
[2nd bit, 1st bit] = [next state, current state]

- 00  dead (next) <- dead (current)
- 01  dead (next) <- live (current)
- 10  live (next) <- dead (current)
- 11  live (next) <- live (current)

To get the current state, simply do

board[i][j] & 1

To get the next state, simply do

board[i][j] >> 1

*/

func gameOfLife(board [][]int) {
	if board == nil || len(board) == 0 {
		return
	}
	m := len(board)
	n := len(board[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			lives := liveNeighbors(board, m, n, i, j)

			if board[i][j] == 1 && lives >= 2 && lives <= 3 {
				board[i][j] = 3
			}
			if board[i][j] == 0 && lives == 3 {
				board[i][j] = 2
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] >>= 1
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func liveNeighbors(board [][]int, m, n, i, j int) int {
	lives := 0
	for x := max(i-1, 0); x <= min(i+1, m-1); x++ {
		for y := max(j-1, 0); y <= min(j+1, n-1); y++ {
			lives += board[x][y] & 1
		}
	}
	lives -= board[i][j] & 1
	return lives
}
