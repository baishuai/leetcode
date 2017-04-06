package p207

/**
There are a total of n courses you have to take, labeled from 0 to n - 1.

Some courses may have prerequisites, for example to take course 0 you have to first take course 1, which is expressed as a pair: [0,1]

Given the total number of courses and a list of prerequisite pairs, is it possible for you to finish all courses?

For example:

2, [[1,0]]
There are a total of 2 courses to take. To take course 1 you should have finished course 0. So it is possible.

2, [[1,0],[0,1]]
There are a total of 2 courses to take. To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1. So it is impossible.
 */

func canFinish(numCourses int, prerequisites [][]int) bool {
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
	for numCourses > 0 {
		if len(in0s) == 0 {
			return false
		}
		in0 := in0s[0]
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
	return true
}
