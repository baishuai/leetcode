package p168

import "bytes"

/**
Given a positive integer, return its corresponding column title as appear in an Excel sheet.

For example:

    1 -> A
    2 -> B
    3 -> C
    ...
    26 -> Z
    27 -> AA
    28 -> AB
 */

func convertToTitle(n int) string {
	var buf bytes.Buffer
	for n > 0 {
		buf.WriteByte(byte((n-1)%26) + 'A')
		n = (n - 1) / 26
	}
	revTitle := buf.Bytes()
	i, j := 0, len(revTitle)-1
	for i < j {
		revTitle[i], revTitle[j] = revTitle[j], revTitle[i]
		i++
		j--
	}
	return string(revTitle)
}
