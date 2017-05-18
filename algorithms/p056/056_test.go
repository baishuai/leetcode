package p056

import "testing"

func test(t *testing.T, intervals []Interval, exp []Interval) {
	ans := merge(intervals)
	if len(ans) != len(exp) {
		t.Fatal("error answer length")
	}
	for i, v := range ans {
		if v.Start != exp[i].Start || v.End != exp[i].End {
			t.Error("error answer", i, v, exp[i])
		}
	}
}

func TestExample(t *testing.T) {
	interval := []Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	exp := []Interval{{1, 6}, {8, 10}, {15, 18}}
	test(t, interval, exp)
}



func TestExampleSort(t *testing.T) {
	interval := []Interval{{1, 3}, {8, 10}, {2, 6}, {15, 18}}
	exp := []Interval{{1, 6}, {8, 10}, {15, 18}}
	test(t, interval, exp)
}



func TestExtra0(t *testing.T) {
	interval := []Interval{{1, 3}, {3,4}}
	exp := []Interval{{1, 4}}
	test(t, interval, exp)
}


func TestExtra1(t *testing.T) {
	interval := []Interval{{1, 2}, {3,4}}
	exp := []Interval{{1, 2},{3,4}}
	test(t, interval, exp)
}

func TestExtra2(t *testing.T) {
	interval := []Interval{{1, 5}, {3,4}}
	exp := []Interval{{1, 5}}
	test(t, interval, exp)
}