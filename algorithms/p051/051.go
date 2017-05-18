package p051

/**
The n-queens puzzle is the problem of placing n queens on an n×n chessboard
 such that no two queens attack each other.

Given an integer n, return all distinct solutions to the n-queens puzzle.

Each solution contains a distinct board configuration of the n-queens' placement,
 where 'Q' and '.' both indicate a queen and an empty space respectively.
 */

/**
For example,
There exist two distinct solutions to the 4-queens puzzle:

[
 [".Q..",  // Solution 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // Solution 2
  "Q...",
  "...Q",
  ".Q.."]
]
 */

func solveNQueens(n int) [][]string {
	locations := make([]int, n)

	res := make([][]string, 0)
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = '.'
	}

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
				oneAns := make([]string, 0)
				for j := 0; j < n; j++ {
					buf[locations[j]] = 'Q'
					oneAns = append(oneAns, string(buf))
					buf[locations[j]] = '.'
				}
				res = append(res, oneAns)
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
