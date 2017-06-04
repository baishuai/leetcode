package p391

import "math"

/**
Given N axis-aligned rectangles where N > 0,
determine if they all together form an exact cover of a rectangular region.

Each rectangle is represented as a bottom-left point and a top-right point.
For example, a unit square is represented as [1,1,2,2].
(coordinate of bottom-left point is (1, 1) and top-right point is (2, 2)).
*/

// 四个角只应出现一次其他点应该出现偶数次
// 使用扫描线方法需要排序，复杂度较高

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type point struct {
	x int
	y int
}

func isRectangleCover(rectangles [][]int) bool {

	if len(rectangles) == 0 {
		return false
	}

	left, bottom, right, top := math.MaxInt32, math.MaxInt32, math.MinInt32, math.MinInt32
	area := 0
	hm := make(map[point]int)
	for _, rect := range rectangles {
		left = min(left, rect[0])
		bottom = min(bottom, rect[1])
		right = max(right, rect[2])
		top = max(top, rect[3])

		area += (rect[2] - rect[0]) * (rect[3] - rect[1])

		lb := point{rect[0], rect[1]}
		lt := point{rect[0], rect[3]}
		rb := point{rect[2], rect[1]}
		rt := point{rect[2], rect[3]}

		hm[lb]++
		hm[lt]++
		hm[rb]++
		hm[rt]++
	}

	lb := point{left, bottom}
	lt := point{left, top}
	rb := point{right, bottom}
	rt := point{right, top}

	if hm[lb] != 1 || hm[lt] != 1 || hm[rb] != 1 || hm[rt] != 1 {
		return false
	} else {
		delete(hm, lb)
		delete(hm, lt)
		delete(hm, rb)
		delete(hm, rt)
	}

	for _, v := range hm {
		if v%2 != 0 {
			return false
		}
	}
	return area == (right-left)*(top-bottom)
}
