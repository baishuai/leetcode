package p416

/**
Given a non-empty array containing only positive integers, find if the array can be partitioned into two subsets such that the sum of elements in both subsets is equal.

Note:
Each of the array element will not exceed 100.
The array size will not exceed 200.
Example 1:

Input: [1, 5, 11, 5]

Output: true

Explanation: The array can be partitioned as [1, 5, 5] and [11].
Example 2:

Input: [1, 2, 3, 5]

Output: false

Explanation: The array cannot be partitioned into equal sum subsets.
**/

func canPartition(nums []int) bool {
	sums := 0
	for _, n := range nums {
		sums += n
	}
	if sums%2 == 1 {
		return false
	}
	sums /= 2

	dp := make([]bool, sums+1)
	dp[0] = true
	for _, num := range nums {
		for i := sums; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}
	}
	return dp[sums]
}
