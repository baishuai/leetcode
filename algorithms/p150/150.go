package p150

/**
Evaluate the value of an arithmetic expression in Reverse Polish Notation.

Valid operators are +, -, *, /. Each operand may be an integer or another expression.

Some examples:
  ["2", "1", "+", "3", "*"] -> ((2 + 1) * 3) -> 9
  ["4", "13", "5", "/", "+"] -> (4 + (13 / 5)) -> 6
 */

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	numStack := make([]int, 0, len(tokens)/2+1)
	for _, t := range tokens {
		switch t {
		case "+":
			b := numStack[len(numStack)-1]
			a := numStack[len(numStack)-2]
			numStack = numStack[:len(numStack)-1]
			numStack[len(numStack)-1] = a + b
		case "-":
			b := numStack[len(numStack)-1]
			a := numStack[len(numStack)-2]
			numStack = numStack[:len(numStack)-1]
			numStack[len(numStack)-1] = a - b
		case "*":
			b := numStack[len(numStack)-1]
			a := numStack[len(numStack)-2]
			numStack = numStack[:len(numStack)-1]
			numStack[len(numStack)-1] = a * b
		case "/":
			b := numStack[len(numStack)-1]
			a := numStack[len(numStack)-2]
			numStack = numStack[:len(numStack)-1]
			numStack[len(numStack)-1] = a / b
		default:
			num, err := strconv.Atoi(t)
			if err == nil {
				numStack = append(numStack, num)
			}
		}
	}
	return numStack[0]
}
