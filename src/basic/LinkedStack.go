package basic

type node struct {
	item interface{}
	next *node
}

type LinkedStack struct {
	size int
	head *node
}

func NewLinkedStack() *LinkedStack {
	return &LinkedStack{0, nil}
}

func (stack *LinkedStack) IsEmpty() bool {
	return stack.size == 0
}

func (stack *LinkedStack) Size() int {
	return stack.size
}

func (stack *LinkedStack) Peek() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	return stack.head.item
}

func (stack *LinkedStack) Push(item interface{}) {
	n := &node{}
	n.item = item
	n.next = stack.head
	stack.head = n
	stack.size++
}

func (stack *LinkedStack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	item := stack.head.item
	stack.head = stack.head.next
	stack.size--
	return item
}
