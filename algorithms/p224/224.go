package p224

/**
Implement a basic calculator to evaluate a simple expression string.

The expression string may contain open ( and closing parentheses ),
 the plus + or minus sign -, non-negative integers and empty spaces .

You may assume that the given expression is always valid.

Some examples:
"1 + 1" = 2
" 2-1 + 2 " = 3
"(1+(4+5+2)-3)+(6+8)" = 23
*/

// + - ( )
func calculate(s string) int {
	ss := []byte(s)

	vStack := make([]int, 0)
	sign := 1
	result, number := 0, 0
	for i := 0; i < len(ss); i++ {
		switch ss[i] {
		case '(':
			vStack = append(vStack, result)
			vStack = append(vStack, sign)
			sign = 1
			result = 0
		case ')':
			result += number * sign
			number = 0
			result *= vStack[len(vStack)-1]
			vStack = vStack[:len(vStack)-1]
			result += vStack[len(vStack)-1]
			vStack = vStack[:len(vStack)-1]
		case '+':
			result += number * sign
			number = 0
			sign = 1
		case '-':
			result += number * sign
			number = 0
			sign = -1
		case ' ':
			continue
		default:
			number = int(ss[i] - '0')
			i++
			for i < len(ss) && ss[i] >= '0' && ss[i] <= '9' {
				number = number*10 + int(ss[i]-'0')
				i++
			}
			if i < len(ss) && ss[i] != ' ' {
				i--
			}
		}
	}
	if number != 0 {
		result += number * sign
	}
	return result
}
