package p171

import "testing"

func test(t *testing.T, title string, exp int) {
	if num := titleToNumber(title); num != exp {
		t.Fatal(num, exp)
	}
}

func TestExample1(t *testing.T) {
	test(t, "A", 1)
}

func TestExample2(t *testing.T) {
	test(t, "B", 2)
}

func TestExample3(t *testing.T) {
	test(t, "AB", 28)
}
