package min_stack

type withMin struct {
	val int
	min int
}

type MinStack struct {
	stack []withMin
}

func Constructor() MinStack {
	return MinStack{stack: []withMin{}}
}

func (m *MinStack) Push(val int) {
	if len(m.stack) == 0 {
		m.stack = append(m.stack, withMin{val: val, min: val})
		return
	}

	min := m.stack[len(m.stack)-1].min
	if val > min {
		m.stack = append(m.stack, withMin{val: val, min: min})
	} else {
		m.stack = append(m.stack, withMin{val: val, min: val})
	}
	return
}

func (m *MinStack) Pop() {
	m.stack = m.stack[:len(m.stack)-1]
}

func (m *MinStack) Top() int {
	return m.stack[len(m.stack)-1].val
}

func (m *MinStack) GetMin() int {
	return m.stack[len(m.stack)-1].min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
