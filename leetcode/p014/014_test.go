package p014

import "testing"

func TestZero(t *testing.T) {
	strs := make([]string, 0)
	ans := longestCommonPrefix(strs)
	if ans != "" {
		t.Fatal("error answer", ans)
	}
}

func TestOne(t *testing.T) {
	strs := []string{"helloworld"}
	ans := longestCommonPrefix(strs)
	if ans != strs[0] {
		t.Fatal("error answer", ans)
	}
}

func TestTwo(t *testing.T)  {
	strs := []string{"qwertyu","qwertyuiop", "qwertgdfdssgf", "qwertfgnjhjh"}
	ans := longestCommonPrefix(strs)
	if ans != "qwert" {
		t.Fatal("error answer", ans)
	}
}