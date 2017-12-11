package p529

func updateBoard(board [][]byte, click []int) [][]byte {

	queue := make([][]int, 0)
	queue = append(queue, click)

	row := len(board)
	if row == 0 {
		return board
	}
	col := len(board[0])

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]

		if board[front[0]][front[1]] == 'M' || board[front[0]][front[1]] == 'X' {
			board[front[0]][front[1]] = 'X'
			continue
		}

		mines := 0
		for i := front[0] - 1; i <= front[0]+1; i++ {
			for j := front[1] - 1; j <= front[1]+1; j++ {
				if i < 0 || i >= row || j < 0 || j >= col {
					continue
				}
				if i == front[0] && j == front[1] {
					continue
				}
				if board[i][j] == 'M' || board[i][j] == 'X' {
					mines++
				}
			}
		}

		if mines > 0 {
			board[front[0]][front[1]] = byte(mines + '0')
		} else {
			board[front[0]][front[1]] = 'B'

			for i := front[0] - 1; i <= front[0]+1; i++ {
				for j := front[1] - 1; j <= front[1]+1; j++ {
					if i < 0 || i >= row || j < 0 || j >= col {
						continue
					}
					if i == front[0] && j == front[1] {
						continue
					}
					if board[i][j] == 'E' {
						board[i][j] = 0
						queue = append(queue, []int{i, j})
					}
				}
			}
		}
	}
	return board
}
