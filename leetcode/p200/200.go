package p200

func numIslands(grid [][]byte) int {
	//func solve(grid [][]byte) {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	id := make([]int, rows*cols+1)
	rank := make([]int, rows*cols+1)
	for i := range id {
		id[i] = i
	}
	count := 0

	find := func(p int) int {
		for p != id[p] {
			id[p] = id[id[p]]
			p = id[p]
		}
		return p
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
		count--
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				count++
				if i == 1 && (grid[i-1][j] == '1') {
					union(i*cols+j, (i-1)*cols+j)
				}
				if j == 1 && (grid[i][j-1] == '1') {
					union(i*cols+j, i*cols+j-1)
				}
				if (i < rows-1 ) && (grid[i+1][j] == '1') {
					union(i*cols+j, (i+1)*cols+j)
				}
				if (j < cols-1) && (grid[i][j+1] == '1') {
					union(i*cols+j, i*cols+j+1)
				}
			}
		}
	}
	return count
}
