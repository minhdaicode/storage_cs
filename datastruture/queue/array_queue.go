package queue

import (
	"bytes"
	"fmt"
)

type ArrayQueue[T comparable] struct {
	front int
	rear  int
	cap   int
	size  int
	arr   []T
}

func NewArrayQueue[T comparable](cap int) *ArrayQueue[T] {
	return &ArrayQueue[T]{
		front: -1,
		rear:  -1,
		cap:   cap,
		size:  0,
		arr:   make([]T, cap),
	}
}

func (q *ArrayQueue[T]) EnQueue(data T) error {
	if q.size == q.cap {
		return ErrIsFull
	}
	q.rear = (q.rear + 1) % q.cap
	q.arr[q.rear] = data
	if q.front == -1 {
		q.front = q.rear
	}
	q.size++
	return nil
}

func (q *ArrayQueue[T]) DeQueue() (T, error) {
	var data T
	if q.IsEmpty() {
		return data, ErrIsEmpty
	}
	data = q.arr[q.front]
	q.front = (q.front + 1) % q.cap
	q.size--
	return data, nil
}

func (q *ArrayQueue[T]) Front() (T, error) {
	var data T
	if q.IsEmpty() {
		return data, ErrIsEmpty
	}
	return q.arr[q.front], nil
}

func (q *ArrayQueue[T]) Rear() (T, error) {
	var data T
	if q.IsEmpty() {
		return data, ErrIsEmpty
	}
	return q.arr[q.rear], nil
}

func (q *ArrayQueue[T]) Size() int {
	return q.size
}

func (q *ArrayQueue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *ArrayQueue[T]) String() string {
	var result bytes.Buffer
	result.WriteByte('[')
	j := q.front
	for i := 0; i < q.size; i++ {
		result.WriteString(fmt.Sprintf("%v", q.arr[j]))
		if i < q.size-1 {
			result.WriteByte(' ')
		}
		j = (j + 1) % q.cap
	}
	result.WriteByte(']')
	return result.String()
}
