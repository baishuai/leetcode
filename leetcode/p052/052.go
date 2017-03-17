package p052

func totalNQueens(n int) int {
	locations := make([]int, n)
	res := 0
	cols := make([]bool, n)
	r45 := make([]bool, n*2-1)
	r135 := make([]bool, n*2-1)

	var solve func(nth int)

	solve = func(nth int) {
		for i := 0; i < n; i++ {
			if cols[i] || r45[n-1-nth+i] || r135[nth+i] {
				continue
			} else {
				cols[i] = true
				r45[n-1-nth+i] = true
				r135[nth+i] = true
				locations[nth] = i
			}

			if nth == n-1 {
				res += 1
			} else {
				solve(nth + 1)
			}
			cols[i] = false
			r45[n-1-nth+i] = false
			r135[nth+i] = false
		}
	}

	solve(0)

	//fmt.Println(res)
	return res
}
