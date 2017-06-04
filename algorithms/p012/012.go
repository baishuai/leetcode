package p012

import "bytes"

/**
Given an integer, convert it to a roman numeral.

Input is guaranteed to be within the range from 1 to 3999.
*/

func intToRoman(num int) string {

	i2sMap := map[int]string{1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX"}
	carryMap := map[byte]byte{'I': 'X', 'V': 'L', 'X': 'C', 'L': 'D', 'C': 'M'}

	var ans bytes.Buffer
	div, pow := 1000, 3
	for div > 0 {
		quo := num / div
		num %= div
		roman := i2sMap[quo]
		for _, char := range []byte(roman) {
			for i := 0; i < pow; i++ {
				char = carryMap[char]
			}
			ans.WriteByte(char)
		}
		div /= 10
		pow--
	}
	return ans.String()
}
