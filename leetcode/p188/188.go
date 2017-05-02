package p188

/**
Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete at most k transactions.

Note:
You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).
 */

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(k int, prices []int) int {

	length := len(prices)
	if length <= 1 {
		return 0
	}

	if (k >= length/2) {
		mProfit := 0
		for i := 1; i < len(prices); i++ {
			diff := prices[i] - prices[i-1]
			if diff > 0 {
				mProfit += diff
			}
		}
		return mProfit
	}

	preDp, dp := make([]int, length), make([]int, length)
	for i := 1; i <= k; i++ {

		preDp, dp = dp, preDp
		buyK := 0 - prices[0]
		dp[0] = 0
		for j := 1; j < length; j++ {
			dp[j] = max(dp[j-1], prices[j]+buyK)
			buyK = max(buyK, preDp[j-1]-prices[j])
		}
	}
	return dp[length-1]
}
