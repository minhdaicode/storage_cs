package linkedlist

type LinkedList[T comparable] interface {
	InsertHead(data T)
	InsertTail(data T)
	Insert(data T, pos int)
	RemoveHead() error
	RemoveTail() error
	Remove(pos int) error
	Tranverse()
	IsEmpty() bool
	Len() int
}
