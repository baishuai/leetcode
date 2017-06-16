package p400

import "strconv"

/**
Find the nth digit of the infinite integer sequence 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...

Note:
n is positive and will fit within the range of a 32-bit signed integer (n < 231).

Example 1:

Input:
3

Output:
3
Example 2:

Input:
11

Output:
0

Explanation:
The 11th digit of the sequence 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ... is a 0, which is part of the number 10.
*/

func findNthDigit(n int) int {
	length := 1
	cnt := 9

	for n > length*cnt {
		n -= length * cnt
		length += 1
		cnt *= 10
	}

	n--
	start := cnt / 9
	start += n / length

	return int(strconv.Itoa(start)[n%length] - '0')
}
