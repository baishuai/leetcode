package p547

//Union Find

func findCircleNum(M [][]int) int {

	count := len(M)
	id := make([]int, count)
	rank := make([]int, count)
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
	size := count
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if M[i][j] == 1 {
				union(i, j)
			}
		}
	}
	return count
}
