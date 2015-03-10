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
	if queue.size+1 > queue.capacity {
		queue.changeCapacity(2 * queue.capacity)
	}
	queue.size++
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
	n := queue.data[queue.first]
	queue.size--
	queue.first++
	if queue.first == queue.capacity {
		queue.first = 0
	}

	if queue.size < queue.capacity/4 {
		queue.changeCapacity(queue.capacity / 4)
	}
	return n
}

func (queue *ArrayResizingQueue) changeCapacity(capacity int) {
	data := queue.data
	queue.data = make([]interface{}, capacity)
	for i := 0; i < queue.size; i++ {
		index := (queue.first + i) % queue.capacity
		queue.data[i] = data[index]
	}
	queue.first = 0
	queue.last = queue.size
	queue.capacity = capacity
}
