package p087

import "testing"

func test(t *testing.T, s1, s2 string, exp bool) {
	if ans := isScramble(s1, s2); ans != exp {
		t.Fatal("error", ans)
	}
}

func TestExample(t *testing.T) {
	test(t, "great", "rgtae", true)
}

func TestExtra(t *testing.T) {
	test(t, "xstjzkfpkggnhjzkpfjoguxvkbuopi", "xbouipkvxugojfpkzjhnggkpfkzjts", true)
}
