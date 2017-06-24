package p539

import (
	"fmt"
	"sort"
)

/**
Given a list of 24-hour clock time points in "Hour:Minutes" format,
 find the minimum minutes difference between any two time points in the list.

Example 1:
Input: ["23:59","00:00"]
Output: 1
Note:
The number of time points in the given list is at least 2 and won't exceed 20000.
The input time is legal and ranges from 00:00 to 23:59.
*/

func minutes(time string) int {
	hour, min := 0, 0
	fmt.Sscanf(time, "%d:%d", &hour, &min)
	return hour*60 + min
}

func findMinDifference(timePoints []string) int {

	mins := make([]int, 0, len(timePoints)+1)
	for _, time := range timePoints {
		mins = append(mins, minutes(time))
	}
	sort.Ints(mins)
	mins = append(mins, 24*60+mins[0])

	res := mins[1] - mins[0]

	for i := 1; i < len(mins); i++ {
		if mins[i]-mins[i-1] < res {
			res = mins[i] - mins[i-1]
		}
	}
	return res
}
