package p354

import (
	"sort"
)

/**
You have a number of envelopes with widths and heights given as a pair of integers (w, h). One envelope can fit into another if and only if both the width and height of one envelope is greater than the width and height of the other envelope.

What is the maximum number of envelopes can you Russian doll? (put one inside other)

Example:
Given envelopes = [[5,4],[6,4],[6,7],[2,3]], the maximum number of envelopes you can Russian doll is 3 ([2,3] => [5,4] => [6,7]).
 */

type Envelope struct {
	w, h int
}

type Envelopes []Envelope

func (es Envelopes) Len() int {
	return len(es)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (es Envelopes) Less(i, j int) bool {
	return (es[i].w < es[j].w) || (es[i].w == es[j].w && es[i].h >= es[j].h)
}

// Swap swaps the elements with indexes i and j.
func (es Envelopes) Swap(i, j int) {
	es[i], es[j] = es[j], es[i]
}

func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	if n == 0 {
		return 0
	}
	ens := make(Envelopes, n)
	for i, e := range envelopes {
		ens[i] = Envelope{e[0], e[1]}
	}
	sort.Sort(ens)
	length := 0
	dp := make([]int, n)
	for _, e := range ens {
		index := sort.SearchInts(dp[:length], e.h)
		dp[index] = e.h
		if index == length {
			length++
		}
	}
	return length
}
