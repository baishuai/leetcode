package p347

import "sort"

/**
For example,
Given [1,1,1,2,2,3] and k = 2, return [1,2].

Note:
You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
Your algorithm's time complexity must be better than O(n log n),
where n is the array's size.

 */

type pair struct {
	n   int
	cnt int
}

type pairs []pair

func (p pairs) Len() int {
	return len(p)
}

func (p pairs) Less(i, j int) bool {
	return p[i].cnt > p[j].cnt
}

func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func topKFrequent(nums []int, k int) []int {

	hm := make(map[int]int)

	for _, n := range nums {
		hm[n]++
	}
	ps := make([]pair, 0, len(hm))

	for k, v := range hm {
		ps = append(ps, pair{k, v})
	}
	sort.Sort(pairs(ps))

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = ps[i].n
	}

	return res
}
