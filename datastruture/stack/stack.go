package stack

type Stack[T comparable] interface {
	Push(data T) error
	Pop() (T, error)
	Peek() (T, error)
	Size() int
	IsEmpty() bool
	IsFull() bool
	Clear()
}
