package p420

/**
A password is considered strong if below conditions are all met:

It has at least 6 characters and at most 20 characters.
It must contain at least one lowercase letter, at least one uppercase letter, and at least one digit.
It must NOT contain three repeating characters in a row ("...aaa..." is weak, but "...aa...a..." is strong, assuming other conditions are met).
Write a function strongPasswordChecker(s), that takes a string s as input, and return the MINIMUM change required to make s a strong password. If s is already strong, return 0.

Insertion, deletion or replace of any one character are all considered as one change.
 */

/**
s.length() < 6
6 <= s.length() <= 20
s.length() > 20
 */

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func strongPasswordChecker(s string) int {
	deleteTarget := max(0, len(s)-20)
	addTarget := max(0, 6-len(s))
	toDelete, toAdd, toReplace := 0, 0, 0
	needUpper, needLower, needDigit := 1, 1, 1

	//len(s) <= 20
	for l, r := 0, 0; r < len(s); r++ {
		b := s[r]
		if b >= 'A' && b <= 'Z' {
			needUpper = 0
		} else if b >= 'a' && b <= 'z' {
			needLower = 0
		} else if b >= '0' && b <= '9' {
			needDigit = 0
		}

		if r-l == 2 {
			if s[l] == s[l+1] && s[l] == s[r] {
				if toAdd < addTarget {
					toAdd++
					l = r
				} else {
					toReplace ++
					l = r + 1
				}
			} else {
				l++
			}
		}

	}

	if len(s) <= 20 {
		return max(addTarget+toReplace, needDigit+needUpper+needLower)
	}

	toReplace = 0
	lenCnts := [3]map[int]int{make(map[int]int), make(map[int]int), make(map[int]int)}
	for l, r := 0, 0; r <= len(s); r++ {
		if r == len(s) || s[l] != s[r] {
			if length := r - l; length > 2 {
				lenCnts[length%3][length]++
			}
			l = r
		}
	}

	for i, numLetters, dec := 0, 0, 0; i < 3; i++ {
		for k, v := range lenCnts[i] {
			if i < 2 {
				numLetters = i + 1
				dec = min(v, (deleteTarget-toDelete)/numLetters)
				toDelete += dec * numLetters
				lenCnts[i][k] -= dec

				if k-numLetters > 2 {
					lenCnts[2][k-numLetters] += dec
				}
			}

			toReplace += lenCnts[i][k] * (k / 3)
		}
	}

	dec := (deleteTarget - toDelete) / 3
	toReplace -= dec
	toDelete -= dec * 3
	return deleteTarget + max(toReplace, needUpper+needLower+needDigit)

}
