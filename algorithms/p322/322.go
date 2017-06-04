package p322

import (
	"math"
)

/**
You are given coins of different denominations and a total amount of money amount. Write a function to compute the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

Example 1:
coins = [1, 2, 5], amount = 11
return 3 (11 = 5 + 5 + 1)

Example 2:
coins = [2], amount = 3
return -1.

Note:
You may assume that you have an infinite number of each kind of coin.
*/

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//DP
func coinChange(coins []int, amount int) int {
	maxCoin := 0
	for _, v := range coins {
		if v > maxCoin {
			maxCoin = v
		}
	}
	if amount == 0 {
		return 0
	} else if maxCoin == 0 {
		return -1
	}
	dpLen := min(maxCoin, amount)

	dp := make([]int, dpLen)
	for i := 1; i <= amount; i++ {
		tmp := math.MaxInt32
		for _, c := range coins {
			ix := i - c
			if ix >= 0 && dp[ix%dpLen] >= 0 {
				tmp = min(tmp, dp[ix%dpLen]+1)
			}
		}
		if tmp == math.MaxInt32 {
			tmp = -1
		}
		dp[i%dpLen] = tmp
	}
	return dp[amount%dpLen]
}
