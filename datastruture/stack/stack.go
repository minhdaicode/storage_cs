package stack

import "errors"

var (
	ErrIsFull  = errors.New("stack is full")
	ErrIsEmpty = errors.New("stack is empty")
)

type Stack[T comparable] interface {
	Push(data T) error
	Pop() (T, error)
	Peek() (T, error)
	Size() int
	IsEmpty() bool
	Clear()
}
