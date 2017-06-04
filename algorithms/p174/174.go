package p174

import "fmt"

/**
The demons had captured the princess (P) and imprisoned her in the bottom-right corner of a dungeon. The dungeon consists of M x N rooms laid out in a 2D grid. Our valiant knight (K) was initially positioned in the top-left room and must fight his way through the dungeon to rescue the princess.

The knight has an initial health point represented by a positive integer. If at any point his health point drops to 0 or below, he dies immediately.

Some of the rooms are guarded by demons, so the knight loses health (negative integers) upon entering these rooms; other rooms are either empty (0's) or contain magic orbs that increase the knight's health (positive integers).

In order to reach the princess as quickly as possible, the knight decides to move only rightward or downward in each step.


Write a function to determine the knight's minimum initial health so that he is able to rescue the princess.

For example, given the dungeon below, the initial health of the knight must be at least 7 if he follows the optimal path RIGHT-> RIGHT -> DOWN -> DOWN.

-2 (K)	-3		3
-5		-10		1
10		30		-5 (P)

Notes:

The knight's health has no upper bound.
Any room can contain threats or power-ups, even the first room the knight enters and the bottom-right room where the princess is imprisoned.
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calculateMinimumHP(dungeon [][]int) int {
	rows := len(dungeon)
	if rows == 0 {
		return 0
	}
	cols := len(dungeon[0])
	preDp := make([]int, cols)
	dp := make([]int, cols)
	for i := rows - 1; i >= 0; i-- {
		dp, preDp = preDp, dp
		for j := cols - 1; j >= 0; j-- {
			if i == rows-1 && j == cols-1 {
				dp[j] = dungeon[i][j]
			} else if i == rows-1 {
				dp[j] = dungeon[i][j] + dp[j+1]
			} else if j == cols-1 {
				dp[j] = dungeon[i][j] + preDp[j]
			} else {
				dp[j] = max(dungeon[i][j]+preDp[j], dungeon[i][j]+dp[j+1])
			}
			if dp[j] >= 0 {
				dp[j] = 0
			}
		}
		fmt.Println(dp)
	}

	return 1 - dp[0]
}
