package p205

import "testing"

func test(t *testing.T, s1, s2 string, exp bool) {
	if ans := isIsomorphic(s1, s2); ans != exp {
		t.Fatal(ans, "Exp", exp)
	}
}

func Test0(t *testing.T) {
	test(t, "egg", "add", true)
}

func Test1(t *testing.T) {
	test(t, "foo", "bar", false)
}

func Test2(t *testing.T) {
	test(t, "ab", "aa", false)
}
