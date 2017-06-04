package p085

/**
Given a 2D binary matrix filled with 0's and 1's, find the largest rectangle
 containing only 1's and return its area.

For example, given the following matrix:

1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0
Return 6.
*/

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	heights := make([]int, len(matrix[0]))
	max := -1
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		if largest := largestRectangleArea(heights); largest > max {
			max = largest
		}
	}

	return max
}

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
