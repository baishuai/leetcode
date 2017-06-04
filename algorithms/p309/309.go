package p309

/**
Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete as many transactions as you like (ie, buy one and sell one share of the stock multiple times) with the following restrictions:

You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).
After you sell your stock, you cannot buy stock on next day. (ie, cooldown 1 day)
Example:

prices = [1, 2, 3, 0, 2]
maxProfit = 3
transactions = [buy, sell, cooldown, buy, sell]
**/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	buy := make([]int, len(prices))
	sell := make([]int, len(prices))
	cool := make([]int, len(prices))
	buy[0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		buy[i] = max(buy[i-1], cool[i-1]-prices[i])
		sell[i] = max(sell[i-1], buy[i-1]+prices[i])
		cool[i] = max(cool[i-1], sell[i-1])
	}
	return sell[len(prices)-1]
}
