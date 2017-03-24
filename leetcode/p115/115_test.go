package p115

import "testing"

func TestExample(t *testing.T) {
	t.Log(numDistinct("rbbbt", "rbbt"))
	t.Log(numDistinct("rbt", "rbbt"))
	t.Log(numDistinct("", "rbbt"))
}

func TestExtra0(t *testing.T) {
	num := numDistinct("aacaacca", "ca")
	if num != 5 {
		t.Fatal(num, 5)
	}
}

func TestExtra1(t *testing.T) {
	num := numDistinct("aacaacca", "cac")
	if num != 4 {
		t.Fatal(num, 4)
	}
}

func TestExtra2(t *testing.T) {
	num := numDistinct("ddd", "dd")
	if num != 3 {
		t.Fatal(num, 3)
	}
}
