package p440

/**
Given integers n and k, find the lexicographically k-th smallest integer in the range from 1 to n.

Note: 1 ≤ k ≤ n ≤ 109.

Example:

Input:
n: 13   k: 2

Output:
10

Explanation:
The lexicographical order is [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9], so the second smallest number is 10.
 */

func calcSteps(n int, ix int) int {
	steps := 0
	nxi := ix + 1
	for ix <= n {
		if n+1 < nxi {
			steps += n + 1 - ix
		} else {
			steps += nxi - ix
		}
		ix *= 10
		nxi *= 10
	}
	return steps
}

func findKthNumber(n int, k int) int {
	ix := 1
	k -= 1
	for k > 0 {
		steps := calcSteps(n, ix)
		if steps <= k {
			ix ++
			k -= steps
		} else {
			ix *= 10
			k -= 1
		}
	}
	return ix
}
