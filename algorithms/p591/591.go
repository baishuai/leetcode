package p591

import "strings"

func isValid(code string) bool {
	stack := make([]string, 0)
	i := 0
	for i < len(code) {
		if i > 0 && len(stack) == 0 {
			return false
		}

		if strings.HasPrefix(code[i:], "<![CDATA[") {
			j := i + 9
			i = strings.Index(code[j:], "]]>") + j
			if i < j {
				return false
			}
			i += 3
		} else if strings.HasPrefix(code[i:], "</") {
			j := i + 2
			i = strings.Index(code[j:], ">") + j
			if i <= j || i-j > 9 {
				return false
			}
			for k := j; k < i; k++ {
				if code[k] > 'Z' || code[k] < 'A' {
					return false
				}
			}
			s := code[j:i]
			i++
			if len(stack) == 0 || s != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else if strings.HasPrefix(code[i:], "<") {
			j := i + 1
			i = strings.Index(code[j:], ">") + j
			if i <= j || i-j > 9 {
				return false
			}
			for k := j; k < i; k++ {
				if code[k] > 'Z' || code[k] < 'A' {
					return false
				}
			}
			s := code[j:i]
			i++
			stack = append(stack, s)
		} else {
			i++
		}
	}
	return len(stack) == 0
}
