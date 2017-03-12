package p020


/**
Given a string containing just the characters '(', ')', '{', '}', '[' and ']',
 determine if the input string is valid.

The brackets must close in the correct order, "()" and "()[]{}" 
 are all valid but "(]" and "([)]" are not.
 */

//use stack

func isValid(s string) bool {
	sLen := len(s)
	if sLen%2 != 0 {
		return false
	}
	//fmt.Println(s)
	hMap := map[byte]byte{')': '(', '}': '{', ']': '['}
	stack := make([]byte, sLen/2)
	stackLen := 0
	for i := 0; i < sLen; i++ {
		v, ok := hMap[s[i]]
		//fmt.Println(i, stack, stackLen)
		if ok {
			if stackLen > 0 && stack[stackLen-1] == v {
				stackLen --
			} else {
				return false
			}
		} else {
			if stackLen >= sLen/2 {
				return false
			} else {
				stack[stackLen] = s[i]
				stackLen++
			}
		}
	}
	return stackLen == 0
}
