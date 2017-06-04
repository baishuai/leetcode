package p084

/**
Given n non-negative integers representing the histogram's bar height where the width of each bar is 1, find the area of largest rectangle in the histogram.


Above is a histogram where width of each bar is 1, given height = [2,1,5,6,2,3].


The largest rectangle is shown in the shaded area, which has area = 10 unit.

For example,
Given heights = [2,1,5,6,2,3],
return 10.
*/

func largestRectangleArea(heights []int) int {
	if heights == nil || len(heights) == 0 {
		return 0
	}
	heightStack := make([]int, 0)
	widthStack := make([]int, 0)
	stackEnd := -1
	largest := -1
	for i := 0; i <= len(heights); i++ {
		v := -1
		if i < len(heights) {
			v = heights[i]
		}
		if v > stackEnd {
			heightStack = append(heightStack, v)
			widthStack = append(widthStack, 1)
			stackEnd = v
		} else if v == stackEnd {
			widthStack[len(widthStack)-1]++
		} else {
			h, w := 0, 0
			for len(heightStack) > 0 && heightStack[len(heightStack)-1] >= v {
				//fmt.Println(heightStack, widthStack)
				h = heightStack[len(heightStack)-1]
				w += widthStack[len(widthStack)-1]
				heightStack = heightStack[:len(heightStack)-1]
				widthStack = widthStack[:len(widthStack)-1]
				if h*w > largest {
					largest = h * w
				}
			}
			heightStack = append(heightStack, v)
			widthStack = append(widthStack, w+1)

			stackEnd = v
		}
	}

	return largest
}

func max(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	}
	return c
}

// too slow
func largestRectangleArea_Slow(heights []int) int {
	if len(heights) == 0 {
		return 0
	} else if len(heights) == 1 {
		return heights[0]
	}
	min := heights[0]
	minIndex := 0
	for i, v := range heights {
		if v < min {
			min = v
			minIndex = i
		}
	}

	return max(len(heights)*min, largestRectangleArea_Slow(heights[:minIndex]), largestRectangleArea_Slow(heights[minIndex+1:]))
}
