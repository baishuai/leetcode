package p123

import "math"

/**
Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete at most two transactions.
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {

	buy1, sold1 := math.MinInt32, 0
	buy2, sold2 := math.MinInt32, 0
	for _, p := range prices {

		sold2 = max(sold2, p+buy2)
		buy2 = max(buy2, sold1-p)
		sold1 = max(sold1, p+buy1)
		buy1 = max(buy1, -p)
	}
	return sold2
}
