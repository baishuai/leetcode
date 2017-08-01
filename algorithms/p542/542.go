package p542

import "math"

/**
Given a matrix consists of 0 and 1, find the distance of the nearest 0 for each cell.

The distance between two adjacent cells is 1.
Example 1:
Input:

0 0 0
0 1 0
0 0 0
Output:
0 0 0
0 1 0
0 0 0
Example 2:
Input:

0 0 0
0 1 0
1 1 1
Output:
0 0 0
0 1 0
1 2 1
Note:
The number of elements of the given matrix will not exceed 10,000.
There are at least one 0 in the given matrix.
The cells are adjacent in only four directions: up, down, left and right.
*/

type pos struct {
	x int
	y int
}

func updateMatrix(matrix [][]int) [][]int {
	if matrix == nil {
		return nil
	}
	queue := make([]pos, 0)
	distance := make([][]int, len(matrix))
	for i := 0; i < len(distance); i++ {
		distance[i] = make([]int, len(matrix[i]))
		copy(distance[i], matrix[i])

		for j := 0; j < len(distance[i]); j++ {
			if distance[i][j] == 0 {
				queue = append(queue, pos{i, j})
			} else {
				distance[i][j] = math.MaxInt32
			}
		}
	}

	dirs := []pos{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(queue) > 0 {
		h := queue[0]
		queue = queue[1:]

		for _, d := range dirs {
			x, y := h.x+d.x, h.y+d.y
			if x >= 0 && x < len(distance) && y >= 0 && y < len(distance[x]) {
				if distance[x][y] > distance[h.x][h.y]+1 {
					distance[x][y] = distance[h.x][h.y] + 1
					queue = append(queue, pos{x, y})
				}
			}
		}
	}
	return distance
}
