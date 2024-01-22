package structures

//======================================== Stack ========================================
type Stack[T any] struct {
	items []T
}

//call func by doing NewStack[type]()
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(i T) {
	s.items = append(s.items, i)
}


func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		var zeroValue T // Default value for T

		return zeroValue
	}

	i := len(s.items) - 1
	item := s.items[i]
	s.items = s.items[0:i]

	return item
}

//======================================== Queue ========================================
type Queue[T any] struct {
	items []T
}

//call func by doing NewQueue[type]()
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(i T) {
	q.items = append(q.items, i)
}

func (q *Queue[T]) Dequeue() T {
	if len(q.items) == 0 {
		var zeroValue T

		return zeroValue
	}

	toRemove := q.items[0]
	q.items = q.items[1:]

	return toRemove
}
