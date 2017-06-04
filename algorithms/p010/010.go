package p010

/**
Implement regular expression matching with support for '.' and '*'.

'.' Matches any single character.
'*' Matches zero or more of the preceding element.

The matching should cover the entire input string (not partial).

The function prototype should be:
bool isMatch(const char *s, const char *p)

Some examples:
isMatch("aa","a") → false
isMatch("aa","aa") → true
isMatch("aaa","aa") → false
isMatch("aa", "a*") → true
isMatch("aa", ".*") → true
isMatch("ab", ".*") → true
isMatch("aab", "c*a*b") → true

*/

func isMatch(s string, p string) bool {
	return match([]byte(s), []byte(p))
}

func match(s []byte, p []byte) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	} else if len(p) == 0 {
		return false
	}

	token := p[0]
	zeroMore := len(p) >= 2 && p[1] == '*'
	if !zeroMore {
		if len(s) > 0 {
			if token != '.' && s[0] != token {
				return false
			} else {
				return match(s[1:], p[1:])
			}
		} else {
			return false
		}
	} else {
		if len(s) > 0 {
			if token != '.' &&
				(len(p) == 2 ||
					(p[2] != token && len(p) > 2 && p[2] != '.' &&
						(len(p) == 3 || (len(p) > 3 && p[3] != '*')))) {
				i := 0
				for i < len(s) && s[i] == token {
					i++
				}
				return match(s[i:], p[2:])
			} else {
				matchLoop := false
				for i := 0; i <= len(s); i++ {
					if token != '.' && i > 0 && s[i-1] != token {
						break
					}
					matchLoop = matchLoop || match(s[i:], p[2:])
				}
				return matchLoop
			}
		} else {
			return match(s, p[2:])
		}
	}
}
