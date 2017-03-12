package p079

func exist(board [][]byte, word string) bool {

	var subExist func(wd []byte, x, y int) bool
	subExist = func(wd []byte, x, y int) bool {
		if len(wd) == 0 {
			return true
		}
		if x < 0 || y < 0 || x >= len(board) || y >= len(board[x]) {
			return false
		}
		if board[x][y] != wd[0] {
			return false
		}
		board[x][y] ^= 255
		wd = wd[1:]
		res := subExist(wd, x, y+1) ||
			subExist(wd, x, y-1) ||
			subExist(wd, x+1, y) ||
			subExist(wd, x-1, y)
		board[x][y] ^= 255
		return res
	}

	ws := []byte(word)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if subExist(ws, i, j) {
				return true
			}
		}
	}
	return false
}
