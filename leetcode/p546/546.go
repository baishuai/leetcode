package p546

/**
Given several boxes with different colors represented by different positive numbers.
You may experience several rounds to remove boxes until there is no box left.
 Each time you can choose some continuous boxes with the same color (composed of k boxes, k >= 1),
 remove them and get k*k points.
Find the maximum points you can get.
 */

type dp struct {
	i, j, k int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func removeBoxes(boxes []int) int {
	memo := make(map[dp]int)
	length := len(boxes)
	init := dp{0, length - 1, 0}
	return dfs(boxes, init, memo)
}

func dfs(boxes []int, pos dp, memo map[dp]int) int {
	if pos.i > pos.j {
		return 0
	} else if memo[pos] > 0 {
		return memo[pos]
	}

	for pos.i < pos.j && boxes[pos.i] == boxes[pos.i+1] {
		pos.i++
		pos.k++
	}

	res := (pos.k+1)*(pos.k+1) + dfs(boxes, dp{pos.i + 1, pos.j, 0}, memo)
	for m := pos.i + 1; m <= pos.j; m++ {
		if boxes[pos.i] == boxes[m] {
			res = max(res, dfs(boxes, dp{pos.i + 1, m - 1, 0}, memo)+
				dfs(boxes, dp{m, pos.j, pos.k + 1}, memo))
		}
	}
	memo[pos] = res
	return res
}
