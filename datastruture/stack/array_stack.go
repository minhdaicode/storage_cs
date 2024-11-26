package stack

import "errors"

var (
	ErrIsFull  = errors.New("stack is full")
	ErrIsEmpty = errors.New("stack is empty")
)

type ArrayStack[T comparable] struct {
	top   int
	cap   uint
	array []T
}

func NewArrayStack[T comparable](cap uint) *ArrayStack[T] {
	return &ArrayStack[T]{
		top:   -1,
		cap:   cap,
		array: make([]T, cap),
	}
}

func (s *ArrayStack[T]) Push(data T) error {
	if s.top == int(s.cap)-1 {
		return ErrIsFull
	}
	s.top++
	s.array[s.top] = data
	return nil
}

func (s *ArrayStack[T]) Pop() (T, error) {
	var data T
	if s.IsEmpty() {
		return data, ErrIsEmpty
	}
	data = s.array[s.top]
	s.top--
	return data, nil
}

func (s *ArrayStack[T]) Peek() (T, error) {
	var data T
	if s.IsEmpty() {
		return data, ErrIsEmpty
	}
	data = s.array[s.top]
	return data, nil
}

func (s *ArrayStack[T]) Size() int {
	return s.top + 1
}

func (s *ArrayStack[T]) IsEmpty() bool {
	return s.top == -1
}

func (s *ArrayStack[T]) Clear() {
	s.array = nil
	s.top = -1
}
