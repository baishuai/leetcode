package p621

import "sort"

func leastInterval(tasks []byte, n int) int {

	tasks26 := make([]int, 26)
	for _, t := range tasks {
		tasks26[t-'A']++
	}

	sort.Ints(tasks26)

	i := 25
	for i >= 0 && tasks26[i] == tasks26[25] {
		i--
	}

	intervals := (tasks26[25]-1)*(n+1) + 25 - i
	if len(tasks) > intervals {
		intervals = len(tasks)
	}
	return intervals
}
