package p068

import (
	"bytes"
)

/**
Given an array of words and a length L, format the text such that each line has exactly L characters and is fully (left and right) justified.

You should pack your words in a greedy approach; that is, pack as many words as you can in each line. Pad extra spaces ' ' when necessary so that each line has exactly L characters.

Extra spaces between words should be distributed as evenly as possible. If the number of spaces on a line do not divide evenly between words, the empty slots on the left will be assigned more spaces than the slots on the right.

For the last line of text, it should be left justified and no extra space is inserted between words.

For example,
words: ["This", "is", "an", "example", "of", "text", "justification."]
L: 16.

Return the formatted lines as:
[
   "This    is    an",
   "example  of text",
   "justification.  "
]
Note: Each word is guaranteed not to exceed L in length.
 */

func fullJustify(words []string, maxWidth int) []string {

	res := make([]string, 0)
	l, h := 0, 0
	buf := bytes.NewBuffer(nil)
	for h < len(words) {
		lenSum := len(words[h])
		for h++; lenSum <= maxWidth && h < len(words); h++ {
			lenSum += 1 + len(words[h])
		}
		//[l,h)
		if lenSum > maxWidth {
			h--
			lenSum -= 1 + len(words[h])
		}
		gap, gapMod := 0, 0
		if h == len(words) {
			gap = 1
		} else if h-l > 1 {
			gap = (maxWidth - (lenSum - (h - l - 1))) / (h - l - 1)
			gapMod = (maxWidth - (lenSum - (h - l - 1))) % (h - l - 1)
		}
		buf.Reset()
		buf.WriteString(words[l])
		for i := 1; i < h-l; i++ {
			for j := 0; j < gap; j++ {
				buf.WriteByte(' ')
			}
			if i <= gapMod {
				buf.WriteByte(' ')
			}
			buf.WriteString(words[i+l])
		}
		for i := buf.Len(); i < maxWidth; i++ {
			buf.WriteByte(' ')
		}
		l = h
		res = append(res, buf.String())
	}
	return res
}
