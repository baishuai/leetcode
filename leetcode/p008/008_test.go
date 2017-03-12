package p008

import "testing"

func TestTrimSpace(t *testing.T) {
	str := "      \f\n1"
	if ans := trimLeadingSpace([]byte(str)); string(ans) != "\f\n1"{
		t.Fatal("error answer", string(ans))
	}

	str = "      "
	if ans := trimLeadingSpace([]byte(str)); string(ans) != ""{
		t.Fatal("error answer", string(ans), ans)
	}
}

func TestBoundary(t *testing.T) {
	str := "21474836471"
	if ans := myAtoi(str); ans != 2147483647{
		t.Fatal("error answer", ans)
	}

	str = "-21474836481"
	if ans := myAtoi(str); ans != -2147483648{
		t.Fatal("error answer", ans)
	}



}


func TestZero(t *testing.T) {
	str := "0"
	if ans := myAtoi(str); ans != 0{
		t.Fatal("error answer", ans)
	}

	str = "-123"
	if ans := myAtoi(str); ans != -123{
		t.Fatal("error answer", ans)
	}


	str = "123"
	if ans := myAtoi(str); ans != 123{
		t.Fatal("error answer", ans)
	}
}


func TestLeadingZero(t *testing.T) {
	str := "-0"
	if ans := myAtoi(str); ans != 0{
		t.Fatal("error answer", ans)
	}

	str = "-00123"
	if ans := myAtoi(str); ans != -123{
		t.Fatal("error answer", ans)
	}


}