package p069

/**
Implement int sqrt(int x).

Compute and return the square root of x.
*/

func mySqrt(x int) int {
	if x < 0 {
		return 0
	}
	s := 0
	for s*s <= x {
		s++
	}
	return s - 1
}
