package linkedlist

import "fmt"

type Circular[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
	size int
}

func NewCircular[T comparable]() *Circular[T] {
	return &Circular[T]{}
}

func (l *Circular[T]) InsertHead(data T) {
	nn := NewNode(data)
	switch {
	case l.IsEmpty():
		l.Head = nn
	case l.size == 1 && l.Head.Next == nil:
		l.Tail = l.Head
		l.Tail.Next = nn
		nn.Next = l.Tail
		l.Head = nn
	default:
		l.Tail.Next = nn
		nn.Next = l.Head
		l.Head = nn
	}
	l.size++
}

func (l *Circular[T]) InsertTail(data T) {
	nn := NewNode(data)
	switch {
	case l.IsEmpty():
		l.Head = nn
	case l.size == 1 && l.Head.Next == nil:
		l.Tail = nn
		l.Tail.Next = l.Head
		l.Head.Next = l.Tail
	default:
		l.Tail.Next = nn
		nn.Next = l.Head
		l.Tail = nn
	}
	l.size++
}

func (l *Circular[T]) Insert(data T, pos int) error {
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

func (l *Circular[T]) RemoveHead() error {
	if l.IsEmpty() {
		return ErrListEmpty
	}
	if l.Head.Next == nil {
		l.Head = nil
	} else {
		l.Head = l.Head.Next
		l.Tail.Next = l.Head

	}
	l.size--
	return nil
}

func (l *Circular[T]) RemoveTail() error {
	if l.IsEmpty() {
		return ErrListEmpty
	}
	switch {
	case l.size == 1 && l.Head.Next == nil:
		l.Head = nil
	case l.size == 2:
		l.Tail = nil
		l.Head.Next = nil
	default:
		cur := l.Head
		prev := cur
		for i := 0; cur.Next != l.Head; i++ {
			prev = cur
			cur = cur.Next
		}
		prev.Next = l.Head
		l.Tail = prev
	}
	l.size--
	return nil
}

func (l *Circular[T]) Remove(pos int) error {
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

func (l *Circular[T]) Tranverse() {
	cur := l.Head
	for cur.Next != nil {
		fmt.Print(cur.Data, " -> ")
		cur = cur.Next
		if cur == l.Tail.Next {
			break
		}
	}
	fmt.Print("head")
}

func (l *Circular[T]) IsEmpty() bool {
	return l.Head == nil || l.size == 0
}

func (l *Circular[T]) Len() int {
	return l.size
}
