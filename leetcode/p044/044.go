package p044


/**
Implement wildcard pattern matching with support for '?' and '*'.

'?' Matches any single character.
'*' Matches any sequence of characters (including the empty sequence).

The matching should cover the entire input string (not partial).

The function prototype should be:
bool isMatch(const char *s, const char *p)

Some examples:
isMatch("aa","a") → false
isMatch("aa","aa") → true
isMatch("aaa","aa") → false
isMatch("aa", "*") → true
isMatch("aa", "a*") → true
isMatch("ab", "?*") → true
isMatch("aab", "c*a*b") → false
 */

func isMatch(s string, p string) bool {

	if (len(s) == 0) && (len(p) == 0) {
		return true
	} else if len(p) == 0 {
		return false
	}

	sb, pb := []byte(s), []byte(p)
	if len(sb) == 0 {
		sb = append(sb, 'a')
		pb = append(pb, 'a')
	}
	dppre := make([]bool, len(pb))
	dpnow := make([]bool, len(pb))

	//first line
	c := sb[0]
	fc := false
	for i, t := range pb {
		if t == '*' {
			dpnow[i] = true
		} else if (t == '?' || t == c) && !fc {
			fc = true
			dpnow[i] = true
		} else {
			break
		}
	}

	for i := 1; i < len(sb); i++ {
		c = sb[i]
		dppre, dpnow = dpnow, dppre
		oneMatch := false
		for j, t := range pb {
			if j == 0 {
				if t == '*' && dppre[j] {
					dpnow[j] = true
					oneMatch = true
				} else {
					dpnow[j] = false
				}
			} else {
				if t == '*' && (dppre[j] || dpnow[j-1]) {
					dpnow[j] = true
					oneMatch = true
				} else if dppre[j-1] && (t == '?' || t == c) {
					dpnow[j] = true
					oneMatch = true
				} else {
					dpnow[j] = false
				}
			}
		}
		if !oneMatch {
			break
		}

	}
	return dpnow[len(pb)-1]
}
