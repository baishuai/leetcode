package p464

/**
In the "100 game," two players take turns adding, to a running total, any integer from 1..10.
The player who first causes the running total to reach or exceed 100 wins.

What if we change the game so that players cannot re-use integers?

For example, two players might take turns drawing from a common pool of numbers of 1..15 without replacement until they reach a total >= 100.

Given an integer maxChoosableInteger and another integer desiredTotal, determine if the first player to move can force a win,
assuming both players play optimally.

You can always assume that maxChoosableInteger will not be larger than 20 and desiredTotal will not be larger than 300.

Example

Input:
maxChoosableInteger = 10
desiredTotal = 11

Output:
false

Explanation:
No matter which integer the first player choose, the first player will lose.
The first player can choose an integer from 1 up to 10.
If the first player choose 1, the second player can only choose integers from 2 up to 10.
The second player will win by choosing 10 and get a total = 11, which is >= desiredTotal.
Same with other integers chosen by the first player, the second player will always win.
*/

//BIT SET

func willWin(bitSet int, desiredTotal int, memo map[int]bool) bool {
	//对方上一步完成后 desiredTotal <= 0， 说明对方已经赢了，己方输了
	if desiredTotal <= 0 {
		return false
	}

	// 如果已经有了计算出来的结果，就直接返回
	if v, ok := memo[bitSet]; ok {
		return v
	}

	for i := uint(1); i <= 20; i++ {
		if bitSet&(1<<i) > 0 {
			bitSetDel := bitSet &^ (1 << i)
			if !willWin(bitSetDel, desiredTotal-int(i), memo) {
				memo[bitSet] = true

				//一旦对手输了，那么确定己方能赢，返回 true
				return true
			}
		}
	}

	// 尝试了所有可选的值之后还不能确定对手会输，那肯定就是对方赢了
	memo[bitSet] = false
	return false
}

func canIWin(maxChoosableInteger int, desiredTotal int) bool {

	sum := (1 + maxChoosableInteger) * maxChoosableInteger / 2
	if sum < desiredTotal {
		return false
	}
	if desiredTotal <= 0 {
		return true
	}

	bitset := 1<<uint(maxChoosableInteger+1) - 2
	memo := make(map[int]bool)
	return willWin(bitset, desiredTotal, memo)
}
