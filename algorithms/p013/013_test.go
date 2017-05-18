package p013

import "testing"

func TestLess10(t *testing.T) {

	if ans := romanToInt("I"); ans != 1 {
		t.Fatal("error answer", ans)
	}

	if ans := romanToInt("II"); ans != 2 {
		t.Fatal("error answer", ans)
	}

	if ans := romanToInt("IV"); ans != 4{
		t.Fatal("error answer", ans)
	}

	if ans := romanToInt("VIII"); ans != 8 {
		t.Fatal("error answer", ans)
	}

	if ans := romanToInt("IX"); ans != 9 {
		t.Fatal("error answer", ans)
	}
}

func Test10b20(t *testing.T) {

	if ans := romanToInt("XI"); ans != 11 {
		t.Fatal("error answer", ans)
	}

	if ans := romanToInt("XII"); ans != 12 {
		t.Fatal("error answer", ans)
	}

	if ans := romanToInt("XV"); ans != 15 {
		t.Fatal("error answer", ans)
	}

	if ans := romanToInt("XVIII"); ans != 18 {
		t.Fatal("error answer", ans)
	}

	if ans := romanToInt("XIX"); ans != 19 {
		t.Fatal("error answer", ans)
	}
}

func TestHundred(t *testing.T) {

	if ans := romanToInt("C"); ans != 100 {
		t.Fatal("error answer", ans)
	}

	if ans := romanToInt("CCC"); ans != 300 {
		t.Fatal("error answer", ans)
	}

	if ans := romanToInt("D"); ans != 500 {
		t.Fatal("error answer", ans)
	}

	if ans := romanToInt("DCCC"); ans != 800 {
		t.Fatal("error answer", ans)
	}

	if ans := romanToInt("M"); ans != 1000 {
		t.Fatal("error answer", ans)
	}
}

func TestOther(t *testing.T)  {
	if ans := romanToInt("DCXXI"); ans != 621 {
		t.Fatal("error answer", ans)
	}
}