package queue

import (
	"bytes"
	"fmt"
)

type SliceQueue[T comparable] struct {
	queue []T
}

func NewSliceQueue[T comparable]() *SliceQueue[T] {
	return &SliceQueue[T]{}
}

func (q *SliceQueue[T]) EnQueue(data T) error {
	q.queue = append(q.queue, data)
	return nil
}

func (q *SliceQueue[T]) DeQueue() (T, error) {
	var data T
	if q.IsEmpty() {
		return data, ErrIsEmpty
	}
	data = q.queue[0]
	q.queue = q.queue[1:]
	return data, nil
}

func (q *SliceQueue[T]) Front() (T, error) {
	var data T
	if q.IsEmpty() {
		return data, ErrIsEmpty
	}
	return q.queue[0], nil
}

func (q *SliceQueue[T]) Rear() (T, error) {
	var data T
	if q.IsEmpty() {
		return data, ErrIsEmpty
	}
	return q.queue[len(q.queue)-1], nil
}

func (q *SliceQueue[T]) Size() int {
	return len(q.queue)
}

func (q *SliceQueue[T]) IsEmpty() bool {
	return len(q.queue) == 0
}

func (q *SliceQueue[T]) String() string {
	var result bytes.Buffer
	result.WriteByte('[')
	for i := 0; i < len(q.queue)-1; i++ {
		result.WriteString(fmt.Sprintf("%v", q.queue[i]))
		if i < len(q.queue)-1 {
			result.WriteByte(' ')
		}
	}
	result.WriteByte(']')
	return result.String()
}
