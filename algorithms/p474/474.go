package p474

/**
In the computer world, use restricted resource you have to generate maximum benefit is what we always want to pursue.

For now, suppose you are a dominator of m 0s and n 1s respectively. On the other hand, there is an array with strings consisting of only 0s and 1s.

Now your task is to find the maximum number of strings that you can form with given m 0s and n 1s. Each 0 and 1 can be used at most once.

Note:
The given numbers of 0s and 1s will both not exceed 100
The size of given string array won't exceed 600.
Example 1:
Input: Array = {"10", "0001", "111001", "1", "0"}, m = 5, n = 3
Output: 4

Explanation: This are totally 4 strings can be formed by the using of 5 0s and 3 1s, which are “10,”0001”,”1”,”0”
Example 2:
Input: Array = {"10", "0", "1"}, m = 1, n = 1
Output: 2

Explanation: You could form "10", but then you'd have nothing left. Better form "0" and "1".
*/

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func findMaxForm(strs []string, m int, n int) int {

	// memo[i][j] 表示i个0，j个1能够装下的最大字符串个数
	memo := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		memo[i] = make([]int, n+1)
	}

	for _, s := range strs {
		num0, num1 := 0, 0
		for _, b := range []byte(s) {
			if b == '1' {
				num1++
			} else {
				num0++
			}
		}

		// 从大到小的顺序不能反，否则结果会出现错误，重复计算了
		for i := m; i >= num0; i-- {
			for j := n; j >= num1; j-- {
				memo[i][j] = max(memo[i][j], memo[i-num0][j-num1]+1)
			}
		}
	}
	return memo[m][n]
}
