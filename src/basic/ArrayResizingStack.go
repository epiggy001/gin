package basic

import "errors"

type ArrayResizingStack struct {
	capacity int
	data     []int
	size     int
}

func NewArrayResizingStack() *ArrayResizingStack {
	return &ArrayResizingStack{
		10,
		make([]int, 10),
		0,
	}
}

func (stack *ArrayResizingStack) IsEmpty() bool {
	return stack.size == 0
}

func (stack *ArrayResizingStack) Peek() (int, error) {
	if stack.IsEmpty() {
		return -1, errors.New("Stack is Empty")
	}
	return stack.data[stack.size-1], nil
}

func (stack *ArrayResizingStack) Push(n int) {
	stack.size += 1
	if stack.size > stack.capacity {
		stack.changeCapacity(stack.capacity * 2)
	}
	stack.data[stack.size-1] = n
	return
}

func (stack *ArrayResizingStack) Pop() (int, error) {
	if stack.IsEmpty() {
		return -1, errors.New("Stack is Empty")
	}
	n := stack.data[stack.size-1]
	stack.size--
	if stack.size < stack.capacity/2 {
		stack.changeCapacity(stack.capacity / 2)
	}
	return n, nil
}

func (stack *ArrayResizingStack) changeCapacity(capacity int) {
	data := make([]int, capacity)
	l := capacity
	if capacity > stack.capacity {
		l = stack.capacity
	}
	for i := 0; i < l; i++ {
		data[i] = stack.data[i]
	}
	stack.data = data
	stack.capacity = capacity
}
