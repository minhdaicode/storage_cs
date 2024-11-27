package queue

import (
	"bytes"
	"fmt"
	"storage_cs/datastruture/linkedlist"
)

type LinkedListQueue[T comparable] struct {
	Head *linkedlist.Node[T]
	Tail *linkedlist.Node[T]
	size int
}

func NewLinkedListQueue[T comparable]() *LinkedListQueue[T] {
	return &LinkedListQueue[T]{}
}

func (q *LinkedListQueue[T]) EnQueue(data T) error {
	nn := linkedlist.NewNode(data)
	if q.IsEmpty() {
		q.Head = nn
	} else {
		cur := q.Tail
		cur.Next = nn
	}
	q.Tail = nn
	q.size++
	return nil
}

func (q *LinkedListQueue[T]) DeQueue() (T, error) {
	var data T
	if q.IsEmpty() {
		return data, ErrIsEmpty
	}
	data = q.Head.Data
	q.Head = q.Head.Next
	q.size--
	return data, nil
}

func (q *LinkedListQueue[T]) Front() (T, error) {
	var data T
	if q.IsEmpty() {
		return data, ErrIsEmpty
	}
	return q.Head.Data, nil
}

func (q *LinkedListQueue[T]) Rear() (T, error) {
	var data T
	if q.IsEmpty() {
		return data, ErrIsEmpty
	}
	return q.Tail.Data, nil
}

func (q *LinkedListQueue[T]) Size() int {
	return q.size
}

func (q *LinkedListQueue[T]) IsEmpty() bool {
	return q.Head == nil
}

func (q *LinkedListQueue[T]) String() string {
	var result bytes.Buffer
	result.WriteByte('[')
	j := q.Head
	for i := 0; i < q.size; i++ {
		result.WriteString(fmt.Sprintf("%v", j.Data))
		if i < q.size-1 {
			result.WriteByte(' ')
		}
		j = j.Next
	}
	result.WriteByte(']')
	return result.String()
}
