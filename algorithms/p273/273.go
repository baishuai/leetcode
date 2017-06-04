package p273

import "bytes"

/**
Convert a non-negative integer to its english words representation.
 Given input is guaranteed to be less than 2^31 - 1.

123 -> "One Hundred Twenty Three"
12345 -> "Twelve Thousand Three Hundred Forty Five"
1234567 -> "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"

*/

func digitToWords(digit int) string {
	word := ""
	if digit >= 0 && digit < 10 {
		word = []string{"Zero", "One", "Two", "Three", "Four",
			"Five", "Six", "Seven", "Eight", "Nine"}[digit]
	}
	return word
}

func tenToWords(digit int) string {
	word := ""
	if digit > 0 && digit < 10 {
		word = []string{"Zero", "Ten", "Twenty", "Thirty", "Forty",
			"Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}[digit]
	}
	return word
}

// convert [10..19] to string
func doubleDigitToWords(digit int) string {
	word := ""
	if digit >= 10 && digit < 20 {
		word = []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen",
			"Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}[digit-10]
	}
	return word
}

func unitNumberToWords(num int) string {
	buf := bytes.Buffer{}
	if num >= 100 {
		buf.WriteString(digitToWords(num / 100))
		buf.WriteByte(' ')
		buf.WriteString("Hundred")
	}
	num %= 100
	if num > 0 {
		if buf.Len() > 0 {
			buf.WriteByte(' ')
		}
		if num < 10 {
			buf.WriteString(digitToWords(num))
			num = 0
		} else if num < 20 {
			buf.WriteString(doubleDigitToWords(num))
			num = 0
		} else {
			buf.WriteString(tenToWords(num / 10))
		}
	}
	num %= 10
	if num > 0 {
		if buf.Len() > 0 {
			buf.WriteByte(' ')
		}
		buf.WriteString(digitToWords(num))
	}
	return buf.String()
}

func numberToWords(num int) string {
	unit := num % 1000
	num /= 1000
	thousand := num % 1000
	num /= 1000
	million := num % 1000
	num /= 1000
	billion := num % 1000

	buf := bytes.Buffer{}
	if billion > 0 {
		buf.WriteString(unitNumberToWords(billion))
		buf.WriteByte(' ')
		buf.WriteString("Billion")
	}
	if million > 0 {
		if buf.Len() > 0 {
			buf.WriteByte(' ')
		}
		buf.WriteString(unitNumberToWords(million))
		buf.WriteByte(' ')
		buf.WriteString("Million")
	}
	if thousand > 0 {
		if buf.Len() > 0 {
			buf.WriteByte(' ')
		}
		buf.WriteString(unitNumberToWords(thousand))
		buf.WriteByte(' ')
		buf.WriteString("Thousand")
	}
	if unit > 0 {
		if buf.Len() > 0 {
			buf.WriteByte(' ')
		}
		buf.WriteString(unitNumberToWords(unit))
	}

	if buf.Len() == 0 {
		buf.WriteString(digitToWords(num))
	}
	return buf.String()
}
