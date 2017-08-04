package p275

import "sort"

func hIndex(citations []int) int {
	sort.Ints(citations)
	h := 1
	for i := len(citations) - 1; i >= 0; i-- {
		if citations[i] < h {
			break
		}
		h++
	}
	return h - 1
}
