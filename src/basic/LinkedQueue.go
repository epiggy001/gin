package basic

type LinkedQueue struct {
	size int
	head *node
	tail *node
}

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{0, nil, nil}
}

func (queue *LinkedQueue) IsEmpty() bool {
	return queue.size == 0
}

func (queue *LinkedQueue) Size() int {
	return queue.size
}

func (queue *LinkedQueue) Enqueue(item interface{}) {
	n := &node{item, nil}
	if queue.tail == nil {
		queue.head = n
	} else {
		queue.tail.next = n
	}
	queue.tail = n
	queue.size++
}

func (queue *LinkedQueue) Dequeue() interface{} {
	if queue.IsEmpty() {
		return nil
	}
	n := queue.head
	queue.head = n.next
	n.next = nil
	if queue.head == nil {
		queue.tail = nil
	}
	queue.size--
	return n.item
}
