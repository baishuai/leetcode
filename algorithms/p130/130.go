package p130

/**
Given a 2D board containing 'X' and 'O' (the letter O), capture all regions surrounded by 'X'.

A region is captured by flipping all 'O's into 'X's in that surrounded region.

For example,
X X X X
X O O X
X X O X
X O X X
After running your function, the board should be:

X X X X
X X X X
X X X X
X O X X
*/

func solve(board [][]byte) {
	rows := len(board)
	if rows <= 2 {
		return
	}
	cols := len(board[0])

	id := make([]int, rows*cols+1)
	rank := make([]int, rows*cols+1)
	for i := range id {
		id[i] = i
	}

	find := func(p int) int {
		for p != id[p] {
			id[p] = id[id[p]]
			p = id[p]
		}
		return p
	}

	connect := func(p, q int) bool {
		return find(p) == find(q)
	}
	union := func(p, q int) {
		p = find(p)
		q = find(q)
		if p == q {
			return
		}
		if rank[p] < rank[q] {
			id[p] = q
		} else if rank[p] > rank[q] {
			id[q] = p
		} else {
			id[q] = p
			rank[p] += 1
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if (i == 0 || i == rows-1 || j == 0 || j == cols-1) && (board[i][j] == 'O') {
				union(i*cols+j, cols*rows)
			} else if board[i][j] == 'O' {

				if i == 1 && (board[i-1][j] == 'O') {
					union(i*cols+j, (i-1)*cols+j)
				}
				if j == 1 && (board[i][j-1] == 'O') {
					union(i*cols+j, i*cols+j-1)
				}
				if board[i+1][j] == 'O' {
					union(i*cols+j, (i+1)*cols+j)
				}
				if board[i][j+1] == 'O' {
					union(i*cols+j, i*cols+j+1)
				}
			}
		}
	}

	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			if board[i][j] == 'O' && !connect(i*cols+j, rows*cols) {
				board[i][j] = 'X'
			}
		}
	}

}
