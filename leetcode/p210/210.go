package p210

// Topological Sort

func findOrder(numCourses int, prerequisites [][]int) []int {

	indegree := make([]int, numCourses)
	in0s := make([]int, 0)
	for _, v := range prerequisites {
		indegree[v[0]] ++
	}

	for i, v := range indegree {
		if v == 0 {
			in0s = append(in0s, i)
		}
	}
	res := make([]int, 0)
	for numCourses > 0 {
		if len(in0s) == 0 {
			return []int{}
		}
		in0 := in0s[0]
		res = append(res, in0)
		in0s = in0s[1:]

		for _, v := range prerequisites {
			if v[1] == in0 {
				indegree[v[0]]--
				if indegree[v[0]] == 0 {
					in0s = append(in0s, v[0])
				}
			}
		}
		numCourses--
	}
	return res
}
