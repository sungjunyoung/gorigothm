package min_stack

import "testing"

func Test_Solution(t *testing.T) {
	m := Constructor()
	m.Push(-2)
	m.Push(0)
	m.Push(-3)
	if m.GetMin() != -3 {
		t.Fatal("failed")
	}
	m.Pop()
	if m.Top() != 0 {
		t.Fatal("failed")
	}
	if m.GetMin() != -2 {
		t.Fatal("failed")
	}

}
