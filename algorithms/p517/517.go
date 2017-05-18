package p517

/**
You have n super washing machines on a line. Initially, each washing machine has some dresses or is empty.

For each move, you could choose any m (1 ≤ m ≤ n) washing machines, and pass one dress of each washing machine to one of its adjacent washing machines at the same time .

Given an integer array representing the number of dresses in each washing machine from left to right on the line, you should find the minimum number of moves to make all the washing machines have the same number of dresses. If it is not possible to do it, return -1.

Example1

Input: [1,0,5]

Output: 3

Explanation:
1st move:    1     0 <-- 5    =>    1     1     4
2nd move:    1 <-- 1 <-- 4    =>    2     1     3
3rd move:    2     1 <-- 3    =>    2     2     2
 */

func max(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	}
	return c
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func findMinMoves(machines []int) int {
	total := 0
	for _, v := range machines {
		total += v
	}
	if total%len(machines) != 0 {
		return -1
	}
	average := total / len(machines)
	moves, cnt := 0, 0
	for _, v := range machines {
		cnt += v - average
		moves = max(moves, abs(cnt), v-average)
	}
	return moves
}
