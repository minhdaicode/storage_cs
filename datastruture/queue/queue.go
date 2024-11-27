package queue

import "errors"

var (
	ErrIsFull  = errors.New("queue is full")
	ErrIsEmpty = errors.New("queue is empty")
)

type Queue[T comparable] interface {
	EnQueue(data T) error
	DeQueue() (T, error)
	Front() (T, error)
	Rear() (T, error)
	Size() int
	IsEmpty() bool
	String() string
}
