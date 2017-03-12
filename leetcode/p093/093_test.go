package p093

import (
	"sort"
	"testing"
)

func test(t *testing.T, s string, exp []string) {
	ans := restoreIpAddresses(s)
	if len(ans) != len(exp) {
		t.Fatal("error answer length")
	}

	sort.Strings(ans)
	for _, v := range exp {
		n := sort.SearchStrings(ans, v)
		if n >= len(ans) || ans[n] != v {
			t.Error("error answer", v, "not exist")
		}
	}
}

func TestExample(t *testing.T) {
	test(t, "25525511135", []string{"255.255.11.135", "255.255.111.35"})
}

func TestExtra0(t *testing.T) {
	test(t, "010010", []string{"0.10.0.10", "0.100.1.0"})
}
