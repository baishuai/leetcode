package p488

/**
Think about Zuma Game. You have a row of balls on the table, colored red(R), yellow(Y), blue(B), green(G), and white(W). You also have several balls in your hand.

Each time, you may choose a ball in your hand, and insert it into the row (including the leftmost place and rightmost place). Then, if there is a group of 3 or more balls in the same color touching, remove these balls. Keep doing this until no more balls can be removed.

Find the minimal balls you have to insert to remove all the balls on the table. If you cannot remove all the balls, output -1.

Examples:

Input: "WRRBBW", "RB"
Output: -1
Explanation: WRRBBW -> WRR[R]BBW -> WBBW -> WBB[B]W -> WW

Input: "WWRRBBWW", "WRBRW"
Output: 2
Explanation: WWRRBBWW -> WWRR[R]BBWW -> WWBBWW -> WWBB[B]WW -> WWWW -> empty

Input:"G", "GGGGG"
Output: 2
Explanation: G -> G[G] -> GG[G] -> empty

Input: "RBYYBBRRB", "YRBGB"
Output: 3
Explanation: RBYYBBRRB -> RBYY[Y]BBRRB -> RBBBRRB -> RRRB -> B -> B[B] -> BB[B] -> empty

Note:
You may assume that the initial row of balls on the table won’t have any 3 or more consecutive balls with the same color.
The number of balls on the table won't exceed 20, and the string represents these balls is called "board" in the input.
The number of balls in your hand won't exceed 5, and the string represents these balls is called "hand" in the input.
Both input strings will be non-empty and only contain characters 'R','Y','B','G','W'.
 */

func findMinStep(board string, hand string) int {

	boards := []byte(board)

	hands := make(map[byte]int)
	for _, v := range []byte(hand) {
		hands[v] ++
	}
	ans := findMinStepHelper(boards, hands)
	if ans < 0 {
		ans = -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMinStepHelper(boards []byte, hands map[byte]int) int {
	length := len(boards)
	if length == 0 {
		return 0
	}
	prev := boards[0]
	prevIndex := 0
	minimun := length << 2
	for i := 1; i <= length; i++ {
		if i < length && boards[i] == prev {
			continue
		} else {
			step := -10
			nextBoards := make([]byte, length-(i-prevIndex))
			copy(nextBoards, boards[:prevIndex])
			copy(nextBoards[prevIndex:], boards[i:])
			if i-prevIndex >= 3 {
				step = findMinStepHelper(nextBoards, hands)
			} else {
				miss := 3 - (i - prevIndex)
				if v, ok := hands[prev]; ok && v >= miss {
					hands[prev] -= miss
					step = miss + findMinStepHelper(nextBoards, hands)
					hands[prev] += miss
				}
			}
			if step >= 0 {
				minimun = min(minimun, step)
			}
			if i < length {
				prev = boards[i]
				prevIndex = i
			}
		}
	}

	if minimun == length<<2 {
		minimun = -10
	}
	return minimun
}
