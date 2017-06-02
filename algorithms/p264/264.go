package p264

/**
Write a program to find the n-th ugly number.

Ugly numbers are positive numbers whose prime factors only include 2, 3, 5. For example, 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 is the sequence of the first 10 ugly numbers.

Note that 1 is typically treated as an ugly number, and n does not exceed 1690.
**/

func nthUglyNumber(n int) int {
	ugly := 1
	q2 := make([]int, 0)
	q3 := make([]int, 0)
	q5 := make([]int, 0)

	n--
	for n > 0 {
		q5 = append(q5, ugly*5)
		if ugly%5 != 0 {
			q3 = append(q3, ugly*3)
			if ugly%3 != 0 {
				q2 = append(q2, ugly*2)
			}
		}

		if q2[0] < q3[0] && q2[0] < q5[0] {
			ugly = q2[0]
			q2 = q2[1:]
		} else if q3[0] < q2[0] && q3[0] < q5[0] {
			ugly = q3[0]
			q3 = q3[1:]
		} else {
			ugly = q5[0]
			q5 = q5[1:]
		}
		n--
	}
	return ugly
}
