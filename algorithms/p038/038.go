package p038

import (
	"bytes"
)

/**
The count-and-say sequence is the sequence of integers beginning as follows:
1, 11, 21, 1211, 111221, ...

1 is read off as "one 1" or 11.
11 is read off as "two 1s" or 21.
21 is read off as "one 2, then one 1" or 1211.
Given an integer n, generate the nth sequence.

Note: The sequence of integers will be represented as a string.
*/

func countAndSay(n int) string {

	var bf bytes.Buffer
	bf.WriteByte('1')
	for n > 1 {
		pre := []byte(bf.String())
		bf.Reset()
		count := byte(0)
		pc := pre[0]
		for _, v := range pre {
			if v == pc {
				count++
			} else {
				bf.WriteByte(count + '0')
				bf.WriteByte(pc)
				pc = v
				count = 1
			}
		}
		bf.WriteByte(count + '0')
		bf.WriteByte(pc)
		n--
	}
	return bf.String()
}
