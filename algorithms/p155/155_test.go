package p155

import "testing"

func TestExample(t *testing.T) {
	s := Constructor()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)

	if s.GetMin() != -3 {
		t.Fail()
	}
	s.Pop()
	if s.Top() != 0 {
		t.Fail()
	}
	if s.GetMin() != -2 {
		t.Fail()
	}
}
