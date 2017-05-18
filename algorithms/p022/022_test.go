package p022

import "testing"

func Test1(t *testing.T)  {
	if ans := generateParenthesis(1); len(ans) != 1{
		t.Fatal("error ans")
	}
}

func Test2(t *testing.T)  {
	if ans := generateParenthesis(2); len(ans) != 2{
		t.Fatal("error ans")
	}
}

func Test3(t *testing.T)  {
	if ans := generateParenthesis(3); len(ans) != 5{
		t.Fatal("error ans")
	}
}

func Test4(t *testing.T)  {
	if ans := generateParenthesis(4); len(ans) != 14{
		t.Fatal("error ans")
	}
}