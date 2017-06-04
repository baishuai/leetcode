package p494

/**
You are given a list of non-negative integers, a1, a2, ..., an, and a target, S. Now you have 2 symbols + and -. For each integer, you should choose one from + and - as its new symbol.

Find out how many ways to assign symbols to make sum of integers equal to target S.

Example 1:
Input: nums is [1, 1, 1, 1, 1], S is 3.
Output: 5
Explanation:

-1+1+1+1+1 = 3
+1-1+1+1+1 = 3
+1+1-1+1+1 = 3
+1+1+1-1+1 = 3
+1+1+1+1-1 = 3

There are 5 ways to assign symbols to make the sum of nums be target 3.
Note:
The length of the given array is positive and will not exceed 20.
The sum of elements in the given array will not exceed 1000.
Your output answer is guaranteed to be fitted in a 32-bit integer.
**/

type entry struct {
	pos int
	sum int
}

func dfs(nums []int, memo map[entry]int, pos int, sum int) int {
	cnt, opos := 0, pos
	if sum < 0 {
		return 0
	} else if v, ok := memo[entry{pos, sum}]; ok {
		return v
	}

	if pos < len(nums) {
		if sum-nums[pos] == 0 {
			cnt++
		}
		cnt += dfs(nums, memo, pos+1, sum-nums[pos])
		pos++
		cnt += dfs(nums, memo, pos, sum)
	}
	memo[entry{opos, sum}] = cnt
	return cnt
}

func findTargetSumWays(nums []int, S int) int {
	sums := 0
	for _, i := range nums {
		sums += i
	}
	if sums < S || (sums-S)%2 == 1 {
		return 0
	}
	delta := (sums - S) / 2

	dpMemo := make(map[entry]int)
	ans := dfs(nums, dpMemo, 0, delta)
	if delta == 0 {
		ans++
	}
	return ans
}
