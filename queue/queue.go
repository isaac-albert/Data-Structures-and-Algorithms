package queue

type node struct {
	next  *node
	Value any
}

type Queue struct {
	head, tail *node
	length     int
}

func (q *Queue) Init() *Queue {
	q.head, q.tail = nil, nil
	q.length = 0
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
	q.length--

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
	q.length++
	if q.tail == nil {
		q.head, q.tail = node, node
		return
	}

	q.tail.next = node
	q.tail = node
}

func (q *Queue) Length() int { return q.length }
