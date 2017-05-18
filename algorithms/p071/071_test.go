package p071

import "testing"

func test(t *testing.T, path, exp string) {
	if ans := simplifyPath(path); ans != exp {
		t.Fatal("error answer", ans, "expected", exp)
	}
}

func TestExample0(t *testing.T) {
	test(t, "/home/", "/home")
}

func TestExample1(t *testing.T) {
	test(t, "/a/./b/../../c/", "/c")
}

func TestExample2(t *testing.T) {
	test(t, "/../", "/")
}

func TestExample3(t *testing.T) {
	test(t, "/", "/")
}

func TestExample4(t *testing.T) {
	test(t, "/home//foo/", "/home/foo")
}
