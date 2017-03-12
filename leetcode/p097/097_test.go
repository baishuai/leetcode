package p097

import "testing"

func test(t *testing.T, s1, s2, s3 string, exp bool) {
	if ans := isInterleave(s1, s2, s3); ans != exp {
		t.Fatal("Error answer", s1, s2, s3)
	}
}

func TestExample0(t *testing.T) {
	test(t, "aabcc", "dbbca", "aadbbcbcac", true)
}

func TestExample1(t *testing.T) {
	test(t, "aabcc", "dbbca", "aadbbbaccc", false)
}


func TestExtra0(t *testing.T) {
	test(t, "aabd", "abdc", "aabdabcd", true)
}

