package p058

/**
Given a string s consists of upper/lower-case alphabets and empty space characters ' ',
 return the length of last word in the string.

If the last word does not exist, return 0.

Note: A word is defined as a character sequence consists of non-space characters only.

For example,
Given s = "Hello World",
return 5.
 */
func lengthOfLastWord(s string) int {
	bss := []byte(s)
	l := len(bss) - 1
	length := 0
	for l >= 0 {
		if (bss[l] >= 'a' && bss[l] <= 'z') ||
			(bss[l] >= 'A' && bss[l] <= 'Z') {
			length ++
			l --
		} else if length == 0 {
			l --
		} else {
			break
		}
	}
	return length
}
