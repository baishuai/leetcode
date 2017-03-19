package p139

import "testing"

func TestWordsByLen(t *testing.T) {
	wordBreak("heoo", []string{"w", "3455", "se", "sd", "r", "drtgds", "vgbh"})
}

func TestExample(t *testing.T) {
	wordBreak("leetcode", []string{"leet", "code"})
}

func test(t *testing.T, words []string, s string, exp bool) {
	if ans := wordBreak(s, words); ans != exp {
		t.Fail()
	}
}

func TestExtra0(t *testing.T) {
	test(t, []string{"leet"}, "leetleet", true)
}

func TestExtra1(t *testing.T) {
	test(t, []string{"leet"}, "", true)
}
