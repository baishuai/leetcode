package p043

/**
Given two non-negative integers num1 and num2 represented as strings,
 return the product of num1 and num2.

Note:

The length of both num1 and num2 is < 110.
Both num1 and num2 contains only digits 0-9.
Both num1 and num2 does not contain any leading zero.
You must not use any built-in BigInteger library
 or convert the inputs to integer directly.
 */

func multiply(num1 string, num2 string) string {
	// 大数乘法
	n1b, n2b := []byte(num1), []byte(num2)
	prod := make([]byte, len(n1b)+len(n2b))
	for i := len(n1b) - 1; i >= 0; i-- {
		//carry := byte(0)
		n1 := n1b[i] - '0'
		for j := len(n2b) - 1; j >= 0; j-- {
			n2 := n2b[j] - '0'
			prod[i+j+1] += (n1 * n2) % 10
			prod[i+j] += (n1*n2)/10 + prod[i+j+1]/10
			prod[i+j+1] %= 10
		}
	}
	for k := 0; k < len(prod); k++ {
		prod[k] += '0'
	}
	firstNoZero := 0
	for firstNoZero < len(prod)-1 {
		if prod[firstNoZero] != '0' {
			break
		}
		firstNoZero++
	}
	return string(prod[firstNoZero:])
}
