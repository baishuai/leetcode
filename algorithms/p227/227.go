package p227

/**
Implement a basic calculator to evaluate a simple expression string.

The expression string contains only non-negative integers, +, -, *, / operators and empty spaces . The integer division should truncate toward zero.

You may assume that the given expression is always valid.

Some examples:
"3+2*2" = 7
" 3/2 " = 1
" 3+5 / 2 " = 5
 */

func calculate(s string) int {
	ss := []byte(s)

	sign := 1
	first, second := 0, 0
	cnt := 0

	var muldiv func(int, int) int

	for i := 0; i < len(ss); i++ {
		switch ss[i] {
		case ' ':
			continue
		case '+':
			sign = 1
		case '-':
			sign = -1
		case '*':
			muldiv = func(a int, b int) int {
				return a * b
			}
		case '/':
			muldiv = func(a int, b int) int {
				return a / b
			}
		default:
			number := int(ss[i] - '0')
			i++
			for i < len(ss) && ss[i] >= '0' && ss[i] <= '9' {
				number = number*10 + int(ss[i]-'0')
				i++
			}
			if i < len(ss) && ss[i] != ' ' {
				i--
			}

			if muldiv != nil {
				if cnt == 1 {
					first = muldiv(first, number)
				} else if cnt == 2 {
					second = muldiv(second, number)
				}
				muldiv = nil
			} else {
				if cnt == 0 {
					first = number * sign
					cnt++
				} else if cnt == 1 {
					second = number * sign
					cnt++
				} else {
					first += second
					second = number * sign
				}

			}
		}
	}
	return first + second
}
