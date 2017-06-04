package p076

import "testing"

func TestExample0(t *testing.T) {
	if ans := minWindow("ADOBECODEBANC", "ABC"); ans != "BANC" {
		t.Error("error anwser", ans)
	}
}

func TestExample1(t *testing.T) {
	if ans := minWindow("A", "A"); ans != "A" {
		t.Error("error anwser", ans)
	}
}

func TestExample2(t *testing.T) {
	if ans := minWindow("Aqwertyuiofhddfg", "Amt"); ans != "" {
		t.Error("error anwser", ans)
	}
}

func TestExample3(t *testing.T) {
	if ans := minWindow("cabwefgewcwaefgcf", "cae"); ans != "cwae" {
		t.Error("error anwser", ans)
	}
}
