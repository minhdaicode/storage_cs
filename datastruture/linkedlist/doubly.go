package linkedlist

import "fmt"

type DLLNode[T comparable] struct {
	Data T
	Prev *DLLNode[T]
	Next *DLLNode[T]
}

func NewDLLNode[T comparable](data T) *DLLNode[T] {
	return &DLLNode[T]{data, nil, nil}
}

type Doubly[T comparable] struct {
	Head *DLLNode[T]
	Tail *DLLNode[T]
	size int
}

func NewDoubly[T comparable]() *Doubly[T] {
	return &Doubly[T]{}
}

func (l *Doubly[T]) InsertHead(data T) {
	nn := NewDLLNode(data)
	switch {
	case l.IsEmpty():
		l.Head = nn
	case l.size == 1 && l.Head.Next == nil:
		l.Tail = l.Head
		l.Tail.Prev = nn
		l.Head = nn
		l.Head.Next = l.Tail
	default:
		l.Head.Prev = nn
		nn.Next = l.Head
		l.Head = nn
	}
	l.size++
}

func (l *Doubly[T]) InsertTail(data T) {
	nn := NewDLLNode(data)
	switch {
	case l.IsEmpty():
		l.Head = nn
	case l.size == 1 && l.Head.Next == nil:
		l.Head.Next = nn
		l.Tail = nn
		l.Tail.Prev = l.Head
	default:
		l.Tail.Next = nn
		nn.Prev = l.Tail
		l.Tail = nn
	}
	l.size++
}

func (l *Doubly[T]) Insert(data T, pos int) error {
	if pos < 0 || pos > l.size {
		return ErrPosOutOfRange
	}
	switch {
	case pos == 0:
		l.InsertHead(data)
	case pos == l.size:
		l.InsertTail(data)
	default:
		cur := l.Head
		prev := cur
		for i := 0; cur.Next != nil; i++ {
			if i == pos {
				nn := NewDLLNode(data)
				prev.Next = nn
				nn.Prev = prev
				nn.Next = cur
				cur.Prev = nn
				l.size++
				break
			}
			prev = cur
			cur = cur.Next
		}
	}
	return nil
}

func (l *Doubly[T]) Tranverse() {
	for cur := l.Head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " -> ")
	}
	fmt.Print("nil")
}

func (l *Doubly[T]) IsEmpty() bool {
	return l.Head == nil || l.size == 0
}

func (l *Doubly[T]) Len() int {
	return l.size
}
