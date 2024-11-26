package stack

type SliceStack[T comparable] struct {
	top   int
	array []T
}

func NewSliceStack[T comparable]() *SliceStack[T] {
	return &SliceStack[T]{
		top:   -1,
		array: make([]T, 0),
	}
}

func (s *SliceStack[T]) Push(data T) error {
	s.array = append(s.array, data)
	s.top++
	return nil
}

func (s *SliceStack[T]) Pop() (T, error) {
	var data T
	if s.IsEmpty() {
		return data, ErrIsEmpty
	}
	data = s.array[s.top]
	s.array = s.array[:s.top]
	s.top--
	return data, nil
}

func (s *SliceStack[T]) Peek() (T, error) {
	var data T
	if s.IsEmpty() {
		return data, ErrIsEmpty
	}
	data = s.array[s.top]
	return data, nil
}

func (s *SliceStack[T]) Size() int {
	return s.top + 1
}

func (s *SliceStack[T]) IsEmpty() bool {
	return s.top == -1
}

func (s *SliceStack[T]) Clear() {
	s.array = nil
	s.top = -1
}
