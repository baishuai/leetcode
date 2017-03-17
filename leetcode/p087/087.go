package p087


func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	} else if len(s1) != len(s2) {
		return false
	}
	return isScrambleBytes([]byte(s1), []byte(s2))
}

func isScrambleBytes(bs1 []byte, bs2 []byte) bool {
	if string(bs1) == string(bs2) {
		return true
	}

	letters := make([]int, 26)
	for i := 0; i < len(bs2); i++ {
		letters[bs1[i]-'a']++
		letters[bs2[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if letters[i] != 0 {
			return false
		}
	}

	for i := 1; i < len(bs2); i++ {
		if isScrambleBytes(bs1[0:i], bs2[0:i]) &&
			isScrambleBytes(bs1[i:], bs2[i:]) {
			return true
		}
		if isScrambleBytes(bs1[0:i], bs2[len(bs2)-i:]) &&
			isScrambleBytes(bs1[i:], bs2[0:len(bs2)-i]) {
			return true
		}
	}
	return false
}
