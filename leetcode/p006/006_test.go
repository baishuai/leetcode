package p006

import "testing"

func Test6Example(t *testing.T) {
	//convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".
	if convert("PAYPALISHIRING", 3) != "PAHNAPLSIIGYIR" {
		t.Fatal("error ans")
	}
}
