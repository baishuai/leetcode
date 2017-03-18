package p125

/**
Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

For example,
"A man, a plan, a canal: Panama" is a palindrome.
"race a car" is not a palindrome.
 */

func isPalindrome(s string) bool {

	bs := []byte(s)

	l, r := 0, len(bs)-1

	for l < r {
		lc := bs[l]
		if lc >= 'A' && lc <= 'Z' {
			lc -= 'A'
		} else if lc >= 'a' && lc <= 'z' {
			lc -= 'a'
		} else if !(lc >= '0' && lc <= '9') {
			l++
			continue
		}
		rc := bs[r]

		if rc >= 'A' && rc <= 'Z' {
			rc -= 'A'
		} else if rc >= 'a' && rc <= 'z' {
			rc -= 'a'
		} else if !(rc >= '0' && rc <= '9') {
			r--
			continue
		}

		if lc == rc {
			l++
			r--
		} else {
			break
		}

	}
	return l >= r
}
