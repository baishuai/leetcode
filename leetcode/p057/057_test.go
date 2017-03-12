package p057

import "testing"

func test(t *testing.T, intervals []Interval, new Interval, exp []Interval) {
	ans := insert(intervals, new)
	if len(ans) != len(exp) {
		t.Fatal("error answer length")
	}
	for i, v := range ans {
		if v.Start != exp[i].Start || v.End != exp[i].End {
			t.Error("error answer", i, v, exp[i])
		}
	}
}

func TestExample0(t *testing.T) {
	interval := []Interval{{1, 3}, {6, 9}}
	new := Interval{2, 5}
	exp := []Interval{{1, 5}, {6, 9}}
	test(t, interval, new, exp)
}

func TestExample1(t *testing.T) {
	interval := []Interval{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	new := Interval{4, 9}
	exp := []Interval{{1, 2}, {3, 10}, {12, 16}}
	test(t, interval, new, exp)
}
