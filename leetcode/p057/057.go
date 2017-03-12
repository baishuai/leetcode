package p057

import "sort"

/**
Given a collection of intervals, merge all overlapping intervals.

For example,
Given [1,3],[2,6],[8,10],[15,18],
return [1,6],[8,10],[15,18].
 */

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

type Interval struct {
	Start int
	End   int
}

type Intervals []Interval

func (is Intervals) Len() int {
	return len(is)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (is Intervals) Less(i, j int) bool {
	return is[i].Start < is[j].Start
}

// Swap swaps the elements with indexes i and j.
func (is Intervals) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	intervals = append(intervals, newInterval)
	sort.Sort(Intervals(intervals))
	if len(intervals) <= 1 {
		return intervals
	}
	ans := make([]Interval, 0)
	merging := intervals[0]
	for _, v := range intervals[1:] {
		if v.Start <= merging.End {
			if v.End > merging.End {
				merging.End = v.End
			}
		} else {
			ans = append(ans, merging)
			merging = v
		}
	}
	ans = append(ans, merging)
	return ans
}
