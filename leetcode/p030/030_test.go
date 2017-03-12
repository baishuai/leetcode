package p030

import "testing"

func TestExample(t *testing.T) {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	ans := findSubstring(s, words)
	t.Log("len", len(ans), ans)
	if len(ans) != 2 {
		t.Fatal("error answer length")
	}
	if ans[0] != 0 || ans[1] != 9 {
		t.Fatal("error answer")
	}
}

func TestCross(t *testing.T) {
	s := "aabbccbbccaadd"
	words := []string{"aa", "bb", "cc", "dd"}
	ans := findSubstring(s, words)
	if len(ans) != 1 {
		t.Fatal("error answer length", len(ans))
	}
	if ans[0] != 6 {
		t.Fatal("error answer")
	}
}

func TestDebug1(t *testing.T) {
	s := "lingmindraboofooowingdingbarrwingmonkeypoundcake"
	words := []string{"fooo", "barr", "wing", "ding", "wing"}
	ans := findSubstring(s, words)
	t.Log("len", len(ans), ans)
	if len(ans) != 1 {
		t.Fatal("error answer length", len(ans))
	}
	if ans[0] != 13 {
		t.Fatal("error answer")
	}
}

func TestDebug2(t *testing.T) {
	s := "barfoofoobarthefoobarman"
	words := []string{"bar", "foo", "the"}
	ans := findSubstring(s, words)
	t.Log("len", len(ans), ans)
	if len(ans) != 3 {
		t.Fatal("error answer length", len(ans))
	}
	if ans[0] != 6 || ans[1] != 9 || ans[2] != 12 {
		t.Fatal("error answer")
	}
}

func TestDebug3(t *testing.T) {
	s := "wordgoodgoodgoodbestword"
	words := []string{"word", "good", "best", "good"}
	ans := findSubstring(s, words)
	t.Log("len", len(ans), ans)
	if len(ans) != 1 {
		t.Fatal("error answer length", len(ans))
	}
	if ans[0] != 8 {
		t.Fatal("error answer")
	}
}


func TestDebug4(t *testing.T) {
	s := "cbaacacbaa"
	words := []string{"cb","aa"}
	ans := findSubstring(s, words)
	t.Log("len", len(ans), ans)
	if len(ans) != 2 {
		t.Fatal("error answer length", len(ans))
	}
	if ans[0] != 0 || ans[1] != 6 {
		t.Fatal("error answer")
	}
}