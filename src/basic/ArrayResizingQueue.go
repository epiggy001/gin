package basic

type ArrayResizingQueue struct {
	capacity int
	data     []interface{}
	size     int
	first    int
	last     int
}

func NewArrayResizingQueue() *ArrayResizingQueue {
	return &ArrayResizingQueue{
		10,
		make([]interface{}, 10),
		0,
		0,
		0,
	}
}

func (queue *ArrayResizingQueue) IsEmpty() bool {
	return queue.size == 0
}

func (queue *ArrayResizingQueue) Peek() interface{} {
	return queue.data[queue.first]
}

func (queue *ArrayResizingQueue) Enqueue(n interface{}) {
	queue.size++
	if queue.size > queue.capacity {
		queue.changeCapacity(2 * queue.capacity)
	}
	queue.data[queue.last] = n
	queue.last++
	if queue.last == queue.capacity {
		queue.last = 0
	}
}

func (queue *ArrayResizingQueue) Dequeue() interface{} {
	if queue.IsEmpty() {
		return nil
	}
	queue.size--
	if queue.size < queue.capacity/4 {
		queue.changeCapacity(queue.capacity / 4)
	}
	n := queue.data[queue.first]
	queue.first++
	if queue.first == queue.capacity {
		queue.first = 0
	}
	return n
}

func (queue *ArrayResizingQueue) changeCapacity(capacity int) {
	data := queue.data
	queue.capacity = capacity
	queue.data = make([]interface{}, capacity)
	count := 0
	for i := queue.first; i < queue.last; i++ {
		queue.data[count] = data[i]
		count++
	}
	queue.first = 0
	queue.last = count

}
