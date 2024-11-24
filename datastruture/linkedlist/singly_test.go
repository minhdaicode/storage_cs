package linkedlist_test

import (
	"storage_cs/datastruture/linkedlist"
	"testing"
)

func TestSingly_InsertHead(t *testing.T) {
	l := linkedlist.NewSingly[int]()
	l.InsertHead(3)
	l.InsertHead(2)
	l.InsertHead(1)
	l.Tranverse()
}

func TestSingly_InsertTail(t *testing.T) {
	l := linkedlist.NewSingly[int]()
	l.InsertTail(3)
	l.InsertTail(2)
	l.InsertTail(1)
	l.Tranverse()
}

func TestSingly_Insert(t *testing.T) {
	l := linkedlist.NewSingly[int]()
	l.Insert(3, 0)
	l.Insert(2, 1)
	l.Insert(1, 2)
	l.Insert(14, 3)
	l.Insert(25, 1)
	l.Tranverse()
}
