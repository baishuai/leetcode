package p685

func root(parent []int, i int) int {
	for i != parent[i] {
		parent[i] = parent[parent[i]]
		i = parent[i]
	}
	return i
}

func findRedundantDirectedConnection(edges [][]int) []int {
	n := len(edges)

	can1 := []int{-1, -1}
	can2 := []int{-1, -1}

	parent := make([]int, n+1)

	for _, edge := range edges {
		if parent[edge[1]] == 0 {
			parent[edge[1]] = edge[0]
		} else {
			can1[0], can1[1] = parent[edge[1]], edge[1]
			can2[0], can2[1] = edge[0], edge[1]
			edge[1] = 0
		}
	}

	for i := 0; i <= n; i++ {
		parent[i] = i
	}

	for _, edge := range edges {
		if edge[1] == 0 {
			continue
		}

		child := edge[1]
		father := edge[0]
		if root(parent, father) == child {
			if can1[0] == -1 {
				return edge
			}
			return can1
		}
		parent[child] = father

	}
	return can2
}
