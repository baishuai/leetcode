package p013

/**
Given a roman numeral, convert it to an integer.

Input is guaranteed to be within the range from 1 to 3999.
*/

func romanToInt(s string) int {
	r2iMap := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	ans := 0
	maybeLeftMinus := -1
	for _, char := range []byte(s) {
		v := r2iMap[char]
		if maybeLeftMinus < v && v/maybeLeftMinus >= 5 && v/maybeLeftMinus <= 10 {
			ans = ans + v - maybeLeftMinus*2
			maybeLeftMinus = -1
		} else {
			maybeLeftMinus = r2iMap[char]
			ans += maybeLeftMinus
		}
	}
	return ans
}
