package p017

/**
Given a digit string, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below.
*/

var (
	i2cMap = map[byte]string{
		'1': "",
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
)

func letterCombinations(digits string) []string {
	return bytesCombinations([]byte(digits))
}

func bytesCombinations(digits []byte) []string {
	ans := make([]string, 0)
	if len(digits) == 0 {
		return ans
	}
	if len(digits) == 1 {
		chars := i2cMap[digits[0]]
		for _, v := range []byte(chars) {
			ans = append(ans, string(v))
		}
		return ans
	} else {
		latters := bytesCombinations(digits[1:])
		chars := i2cMap[digits[0]]
		for _, v := range []byte(chars) {
			//ans = append(ans, string(v))
			for _, latter := range latters {
				ans = append(ans, string(v)+latter)
			}
		}
		return ans
	}
}
