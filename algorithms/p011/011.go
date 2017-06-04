package p011

/**
Given n non-negative integers a1, a2, ..., an, where each represents a point at coordinate (i, ai).
n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0).
Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.
*/

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	max := 0
	for left != right {
		leftV := height[left]
		rightV := height[right]
		//fmt.Println(left, right, leftV, rightV)
		minHeight := 0
		if leftV < rightV {
			left++
			minHeight = leftV
		} else {
			right--
			minHeight = rightV
		}
		area := minHeight * (1 + right - left)
		//fmt.Println("Area", area)
		if area > max {
			max = area
		}
	}
	return max
}
