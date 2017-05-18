package p068

import (
	"testing"
)

//["this","is","an","exampleee"]
//17
//["this     is    an","exampleee        "]

func test(t *testing.T, words []string, max int, exp []string) {
	res := fullJustify(words, max)
	if len(res) != len(exp) {
		t.Fatal("error answer length")
	}
	for i, v := range exp {
		if res[i] != v {
			t.Error("error value", i, res[i], v)
		}
	}
}

func TestEXample0(t *testing.T) {
	test(t, []string{"this", "is"}, 6, []string{"this  ", "is    "})
}

func TestExample1(t *testing.T) {
	test(t, []string{"this", "is", "an", "exampleee"}, 17,
		[]string{"this     is    an", "exampleee        "})
}

func TestExample2(t *testing.T) {
	test(t, []string{"This", "is", "an", "example", "of", "text", "justification."}, 16,
		[]string{"This    is    an", "example  of text", "justification.  "})
}

func TestExample3(t *testing.T) {
	test(t, []string{"What", "must", "be", "shall", "be."}, 12, []string{"What must be", "shall be.   "})
}
