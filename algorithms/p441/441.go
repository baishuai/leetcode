package p441

import "sort"

/**
You have a total of n coins that you want to form in a staircase shape, where every k-th row must have exactly k coins.

Given n, find the total number of full staircase rows that can be formed.

n is a non-negative integer and fits within the range of a 32-bit signed integer.

Example 1:

n = 5

The coins can form the following rows:
¤
¤ ¤
¤ ¤

Because the 3rd row is incomplete, we return 2.
Example 2:

n = 8

The coins can form the following rows:
¤
¤ ¤
¤ ¤ ¤
¤ ¤
*/

// 求解数学不等式

func arrangeCoins(n int) int {

	// x * (x+1) / 2  <= n
	// x * (x+1) <= n*2
	return sort.Search(n+1, func(i int) bool {
		return i*(i+1)/2 > n
	}) - 1
}
