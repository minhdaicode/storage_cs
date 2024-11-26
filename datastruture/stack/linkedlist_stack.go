package stack

import "storage_cs/datastruture/linkedlist"

type LinkedListStack[T comparable] struct {
	top  *linkedlist.Node[T]
	size int
}

func NewLinkedListStack[T comparable]() *LinkedListStack[T] {
	return &LinkedListStack[T]{}
}

func (s *LinkedListStack[T]) Push(data T) error {
	s.top = &linkedlist.Node[T]{Data: data, Next: s.top}
	s.size++
	return nil
}

func (s *LinkedListStack[T]) Pop() (T, error) {
	var data T
	if s.IsEmpty() {
		return data, ErrIsEmpty
	}
	data, s.top = s.top.Data, s.top.Next
	return data, nil
}

func (s *LinkedListStack[T]) Peek() (T, error) {
	var data T
	if s.IsEmpty() {
		return data, ErrIsEmpty
	}
	data = s.top.Data
	return data, nil
}

func (s *LinkedListStack[T]) Size() int {
	return s.size
}

func (s *LinkedListStack[T]) IsEmpty() bool {
	return s.size == 0
}

func (s *LinkedListStack[T]) Clear() {
	s.top = nil
	s.size = 0
}
