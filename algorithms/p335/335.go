package p335

/**
You are given an array x of n positive numbers. You start at point (0,0) and moves x[0] metres to the north,
 then x[1] metres to the west, x[2] metres to the south, x[3] metres to the east and so on.
 In other words, after each move your direction changes counter-clockwise.

Write a one-pass algorithm with O(1) extra space to determine, if your path crosses itself, or not.
*/

func isSelfCrossing(x []int) bool {
	if len(x) < 4 {
		return false
	}

	trapLo, trapHi, trapIn := 0, 0, -1
	for i := 3; i < len(x); i++ {
		if x[i-1] <= x[i-3] && x[i] >= x[i-2] {
			return true
		}
		if trapIn >= 0 {
			if x[i] >= trapIn {
				return true
			} else {
				trapLo, trapHi, trapIn = 0, 0, -1
			}
		}
		if trapLo > 0 {
			if x[i-1] == x[i-3] && x[i] >= trapLo {
				return true
			}
			if x[i] >= trapLo && x[i] <= trapHi {
				trapIn = x[i-1] - x[i-3]
			} else if x[i] < trapIn {
				trapLo, trapHi = 0, 0
			} else {
				trapLo = x[i-1] - x[i-3]
				trapHi = x[i-1]
			}
		}
		if x[i-1] >= x[i-3] && x[i] >= x[i-2] {
			trapLo = x[i-1] - x[i-3]
			trapHi = x[i-1]
		}
	}
	return false
}
