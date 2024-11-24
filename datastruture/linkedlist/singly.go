package linkedlist

import (
	"errors"
	"fmt"
)

var (
	ErrListEmpty     = errors.New("list is empty")
	ErrPosOutOfRange = errors.New("position out of range")
)

type Node[T comparable] struct {
	Data T
	Next *Node[T]
}

func NewNode[T comparable](data T) *Node[T] {
	return &Node[T]{data, nil}
}

type Singly[T comparable] struct {
	Head *Node[T]
	size int
}

func NewSingly[T comparable]() *Singly[T] {
	return &Singly[T]{}
}

func (l *Singly[T]) InsertHead(data T) {
	nn := NewNode(data)
	nn.Next = l.Head
	l.Head = nn
	l.size++
}

func (l *Singly[T]) InsertTail(data T) {
	if l.IsEmpty() {
		l.InsertHead(data)
		return
	}
	cur := l.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = NewNode(data)
	l.size++
}

func (l *Singly[T]) Insert(data T, pos int) error {
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
				nn := NewNode(data)
				prev.Next = nn
				nn.Next = cur
				l.size++
				break
			}
			prev = cur
			cur = cur.Next
		}
	}
	return nil
}

func (l *Singly[T]) RemoveHead() error {
	if l.IsEmpty() {
		return ErrListEmpty
	}
	l.Head = l.Head.Next
	l.size--
	return nil
}

func (l *Singly[T]) RemoveTail() error {
	switch {
	case l.IsEmpty():
		return ErrListEmpty
	case l.Head.Next == nil:
		return l.RemoveHead()
	default:
		cur := l.Head
		for cur.Next.Next != nil {
			cur = cur.Next
		}
		cur.Next = nil
		l.size--
		return nil
	}
}

func (l *Singly[T]) Remove(pos int) error {
	switch {
	case pos < 0 || pos > l.size:
		return ErrPosOutOfRange
	case pos == 0:
		return l.RemoveHead()
	case pos == l.size:
		return l.RemoveTail()
	default:
		cur := l.Head
		prev := cur
		for i := 0; cur.Next != nil; i++ {
			if i == pos {
				prev.Next = cur.Next
				l.size--
				break
			}
			prev = cur
			cur = cur.Next
		}
		return nil
	}
}

func (l *Singly[T]) Tranverse() {
	for cur := l.Head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " -> ")
	}
	fmt.Print("nil")
}

func (l *Singly[T]) IsEmpty() bool {
	return l.Head == nil || l.size == 0
}

func (l *Singly[T]) Len() int {
	return l.size
}
