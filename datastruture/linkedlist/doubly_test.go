package linkedlist_test

import (
	"fmt"
	"reflect"
	"storage_cs/datastruture/linkedlist"
	"testing"
)

func TestDoubly_InsertHead(t *testing.T) {
	l := linkedlist.NewDoubly[int]()
	l.InsertHead(3)
	l.InsertHead(2)
	l.InsertHead(1)

	want := []int{1, 2, 3}
	got := []int{}
	cur := l.Head
	got = append(got, cur.Data)

	for cur.Next != nil {
		cur = cur.Next
		got = append(got, cur.Data)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v , want: %v", got, want)
	}
}

func TestDoubly_InsertTail(t *testing.T) {
	l := linkedlist.NewDoubly[int]()
	l.InsertTail(3)
	l.InsertTail(2)
	l.InsertTail(1)

	want := []int{3, 2, 1}
	got := []int{}
	cur := l.Head
	got = append(got, cur.Data)

	for cur.Next != nil {
		cur = cur.Next
		got = append(got, cur.Data)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v , want: %v", got, want)
	}
}

func TestDoubly_Insert(t *testing.T) {
	l := linkedlist.NewDoubly[int]()

	err := l.Insert(12, -2)
	if err != linkedlist.ErrPosOutOfRange {
		t.Errorf("something went wrong")
	}

	l.Insert(3, 0)
	l.Insert(2, 1)
	l.Insert(1, 2)
	l.Insert(14, 3)
	l.Insert(25, 1)

	want := []int{3, 25, 2, 1, 14}
	got := []int{}
	cur := l.Head
	got = append(got, cur.Data)

	for cur.Next != nil {
		cur = cur.Next
		got = append(got, cur.Data)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v , want: %v", got, want)
	}

	err = l.Insert(12, 7)
	if err != linkedlist.ErrPosOutOfRange {
		t.Errorf("wsomething went wrong")
	}

	l.Tranverse()
	fmt.Println()
	fmt.Printf("size of list: %v", l.Len())
}
