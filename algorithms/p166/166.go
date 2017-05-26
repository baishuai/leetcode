package p166

import (
	"bytes"
	"strconv"
)

/**
Given two integers representing the numerator and denominator of a fraction, return the fraction in string format.

If the fractional part is repeating, enclose the repeating part in parentheses.

For example,

Given numerator = 1, denominator = 2, return "0.5".
Given numerator = 2, denominator = 1, return "2".
Given numerator = 2, denominator = 3, return "0.(6)".
**/

// 按照小学生时候的那种长除法来计算
// 注意处理正负号的问题，都是用正数计算，前面加上符号就好
// 不必担心会出现无线不循环小数，因为无线不循环小数就是无理数了，而这里计算的都是有理数
// 答案共分为三部分，小数点之前的部分，紧跟小数点后不循环的部分，可能的小数点最后循环的部分

func fractionToDecimal(numerator int, denominator int) string {
	buf := bytes.Buffer{}
	if numerator < 0 && denominator > 0 {
		numerator = -numerator
		buf.WriteByte('-')
	}
	if numerator > 0 && denominator < 0 {
		denominator = -denominator
		buf.WriteByte('-')
	}

	integer := numerator / denominator

	numerator = numerator % denominator
	decimals := make([]byte, 0)
	reminders := make(map[int]int)
	var loop string
	for numerator != 0 {
		if v, ok := reminders[numerator]; ok {
			//loop
			loop = string(decimals[v:])
			decimals = decimals[:v]
			break
		}
		reminders[numerator] = len(decimals)
		numerator *= 10
		decimals = append(decimals, byte(numerator/denominator)+'0')
		numerator %= denominator
	}

	buf.WriteString(strconv.Itoa(integer))
	if len(decimals) > 0 || len(loop) > 0 {
		buf.WriteByte('.')
	}
	buf.WriteString(string(decimals))
	if len(loop) > 0 {
		buf.WriteByte('(')
		buf.WriteString(loop)
		buf.WriteByte(')')
	}
	return buf.String()
}
