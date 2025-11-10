package queue

type node struct {
	next  *node
	Value any
}

type Queue struct {
	head, tail *node
	Length     int
}

func (q *Queue) Init() *Queue {
	q.head, q.tail = nil, nil
	q.Length = 0
	return q
}

func New() *Queue {
	return new(Queue).Init()
}

func (q *Queue) Peek() any {
	if q.head != nil {
		return q.head.Value
	}

	return nil
}

func (q *Queue) Dequeue() any {
	if q.head == nil {
		return nil
	}
	q.Length--

	h := q.head
	q.head = q.head.next

	if q.head == nil {
		q.tail = nil
	}

	//free head
	h.next = nil

	return h.Value
}

func (q *Queue) Enqueue(item any) {
	node := &node{
		Value: item,
	}
	q.Length++
	if q.tail == nil {
		q.head, q.tail = node, node
		return
	}

	q.tail.next = node
	q.tail = node
}

