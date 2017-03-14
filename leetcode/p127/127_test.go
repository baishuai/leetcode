package p127

import "testing"

func TestExample(t *testing.T) {
	if ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}) != 5 {
		t.Fatal("error")
	}
}
