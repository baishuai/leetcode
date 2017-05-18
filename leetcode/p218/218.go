package p218

import (
	"sort"
)

/**
218. The Skyline Problem
https://leetcode.com/problems/the-skyline-problem/#/description

A city's skyline is the outer contour of the silhouette formed by all the buildings in that city when viewed from a distance. Now suppose you are given the locations and height of all the buildings as shown on a cityscape photo (Figure A), write a program to output the skyline formed by these buildings collectively (Figure B).
**/

type building struct {
	l, r, h int
}

func criticalPos(buildings [][]int) []building {
	res := make(map[int]struct{})
	empty := struct{}{}
	for _, b := range buildings {
		res[b[0]] = empty
		res[b[1]] = empty
	}
	criticals := make([]int, 0, len(res)+1)
	for v := range res {
		criticals = append(criticals, v)
	}
	sort.Ints(criticals)
	ans := make([]building, len(criticals))
	criticals = append(criticals, criticals[len(criticals)-1]+1)
	for i := 0; i < len(ans); i++ {
		ans[i] = building{criticals[i], criticals[i+1], 0}
	}
	return ans
}

func getSkyline(buildings [][]int) [][]int {
	// get critical point(x position)
	// [Li, Ri, Hi]
	if len(buildings) == 0 {
		return buildings
	}
	criticals := criticalPos(buildings)

	for _, b := range buildings {
		idx := sort.Search(len(criticals), func(i int) bool {
			return criticals[i].l >= b[0]
		})
		for idx < len(criticals) && criticals[idx].l >= b[0] &&
			criticals[idx].r <= b[1] {
			if b[2] > criticals[idx].h {
				criticals[idx].h = b[2]
			}
			idx++
		}
	}
	ans := make([][]int, 0, len(criticals))
	preHeight := -1
	for _, b := range criticals {
		if b.h == preHeight {
			continue
		}
		ans = append(ans, []int{b.l, b.h})
		preHeight = b.h
	}
	return ans
}
