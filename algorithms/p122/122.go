package p122

/**
Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete as many transactions as you like
 (ie, buy one and sell one share of the stock multiple times).
  However, you may not engage in multiple transactions at the same time
   (ie, you must sell the stock before you buy again).
*/

func maxProfit(prices []int) int {
	mProfit := 0
	for i := 1; i < len(prices); i++ {
		diff := prices[i] - prices[i-1]
		if diff > 0 {
			mProfit += diff
		}
	}
	return mProfit
}
