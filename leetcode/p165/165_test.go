package p165

import "testing"

func test(t *testing.T, v1, v2 string, exp int) {
	ans := compareVersion(v1, v2)
	if ans != exp {
		t.Fatal("error", ans, "exp", exp)
	}
}

func TestExample(t *testing.T) {
	test(t, ".1", ".2", -1)
}
