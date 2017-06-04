package p005

import "testing"

func TestExample(t *testing.T) {
	if ans := longestPalindrome("babad"); ans != "aba" && ans != "bab" {
		t.Fatal("error ans", ans)
	}

	if ans := longestPalindrome("cbbd"); ans != "bb" {
		t.Fatal("error ans", ans)
	}
}

func TestAba(t *testing.T) {
	if ans := longestPalindrome("aba"); ans != "aba" {
		t.Fatal("error ans", ans)
	}

	if ans := longestPalindrome("babab"); ans != "babab" {
		t.Fatal("error ans", ans)
	}
}
