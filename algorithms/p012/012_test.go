package p012

import "testing"

func TestLess10(t *testing.T) {

	if ans := intToRoman(1); ans != "I" {
		t.Fatal("error answer", ans)
	}

	if ans := intToRoman(2); ans != "II" {
		t.Fatal("error answer", ans)
	}

	if ans := intToRoman(5); ans != "V" {
		t.Fatal("error answer", ans)
	}

	if ans := intToRoman(8); ans != "VIII" {
		t.Fatal("error answer", ans)
	}

	if ans := intToRoman(9); ans != "IX" {
		t.Fatal("error answer", ans)
	}
}

func Test10b20(t *testing.T) {

	if ans := intToRoman(11); ans != "XI" {
		t.Fatal("error answer", ans)
	}

	if ans := intToRoman(12); ans != "XII" {
		t.Fatal("error answer", ans)
	}

	if ans := intToRoman(15); ans != "XV" {
		t.Fatal("error answer", ans)
	}

	if ans := intToRoman(18); ans != "XVIII" {
		t.Fatal("error answer", ans)
	}

	if ans := intToRoman(19); ans != "XIX" {
		t.Fatal("error answer", ans)
	}
}

func TestHundred(t *testing.T) {

	if ans := intToRoman(100); ans != "C" {
		t.Fatal("error answer", ans)
	}

	if ans := intToRoman(300); ans != "CCC" {
		t.Fatal("error answer", ans)
	}

	if ans := intToRoman(500); ans != "D" {
		t.Fatal("error answer", ans)
	}

	if ans := intToRoman(800); ans != "DCCC" {
		t.Fatal("error answer", ans)
	}

	if ans := intToRoman(1000); ans != "M" {
		t.Fatal("error answer", ans)
	}
}
