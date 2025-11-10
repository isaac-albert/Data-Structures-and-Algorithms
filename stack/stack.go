package stack

type node struct {
	Value any
	prev  *node
}

type Stack struct {
	head   *node
	Length int
}

func (s *Stack) Init() *Stack {
	s.head = nil
	s.Length = 0
	return s
}

func New() *Stack { return new(Stack).Init() }

func (s *Stack) Push(item any) {

	node := &node{Value: item}
	s.Length++
	if s.head == nil {
		s.head = node
		return
	}

	node.prev = s.head
	s.head = node
}

func (s *Stack) Pop() any {

	if s.head == nil {
		return nil
	}

	s.Length--
	h := s.head
	s.head = s.head.prev

	if s.Length == 0 {
		s.head = nil
	}

	//free space
	h.prev = nil

	return h.Value
}

func (s *Stack) Peek() any {
	if s.head == nil {
		return nil
	}

	return s.head.Value
}
