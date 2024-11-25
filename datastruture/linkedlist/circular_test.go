package linkedlist_test

import (
	"fmt"
	"reflect"
	"storage_cs/datastruture/linkedlist"
	"testing"
)

func TestCircular_InsertHead(t *testing.T) {
	l := linkedlist.NewCircular[int]()
	l.InsertHead(3)
	l.InsertHead(2)
	l.InsertHead(1)

	want := []int{1, 2, 3}
	got := []int{}
	cur := l.Head
	got = append(got, cur.Data)

	for cur.Next != nil {
		cur = cur.Next
		if cur == l.Tail.Next {
			break
		}
		got = append(got, cur.Data)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v , want: %v", got, want)
	}
}

func TestCircular_InsertTail(t *testing.T) {
	l := linkedlist.NewCircular[int]()
	l.InsertTail(3)
	l.InsertTail(2)
	l.InsertTail(1)

	want := []int{3, 2, 1}
	got := []int{}
	cur := l.Head
	got = append(got, cur.Data)

	for cur.Next != nil {
		cur = cur.Next
		if cur == l.Tail.Next {
			break
		}
		got = append(got, cur.Data)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v , want: %v", got, want)
	}
}

func TestCircular_Insert(t *testing.T) {
	l := linkedlist.NewCircular[int]()

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
		if cur == l.Tail.Next {
			break
		}
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

func TestCircular_RemoveHead(t *testing.T) {
	l := linkedlist.NewCircular[int]()

	err := l.RemoveHead()
	if err != linkedlist.ErrListEmpty {
		t.Errorf("something went wrong")
	}

	l.InsertHead(3)
	l.InsertHead(2)
	l.InsertHead(1)

	err = l.RemoveHead()
	if err != nil {
		t.Errorf("something went wrong")
	}

	want := []int{2, 3}
	got := []int{}
	cur := l.Head
	got = append(got, cur.Data)

	for cur.Next != nil {
		cur = cur.Next
		if cur == l.Tail.Next {
			break
		}
		got = append(got, cur.Data)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v , want: %v", got, want)
	}
}

func TestCircular_RemoveTail(t *testing.T) {
	l := linkedlist.NewCircular[int]()

	err := l.RemoveTail()
	if err != linkedlist.ErrListEmpty {
		t.Errorf("something went wrong")
	}

	l.InsertHead(3)

	err = l.RemoveTail()
	if err != nil {
		t.Errorf("something went wrong")
	}

	l.InsertTail(7)
	l.InsertHead(76)

	err = l.RemoveTail()
	if err != nil {
		t.Errorf("something went wrong")
	}

	l.InsertHead(2)
	l.InsertHead(1)
	l.InsertTail(12)
	l.InsertTail(5)

	err = l.RemoveTail()
	if err != nil {
		t.Errorf("something went wrong")
	}

	want := []int{1, 2, 76, 12}
	got := []int{}
	cur := l.Head
	got = append(got, cur.Data)

	for cur.Next != nil {
		cur = cur.Next
		if cur == l.Tail.Next {
			break
		}
		got = append(got, cur.Data)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v , want: %v", got, want)
	}
}

func TestCircular_Remove(t *testing.T) {
	l := linkedlist.NewCircular[int]()

	err := l.Remove(-2)
	if err != linkedlist.ErrPosOutOfRange {
		t.Errorf("something went wrong")
	}

	l.InsertHead(3)

	err = l.Remove(0)
	if err != nil {
		t.Errorf("something went wrong")
	}

	l.InsertHead(2)
	l.InsertHead(1)
	l.InsertTail(12)
	l.InsertTail(5)

	err = l.Remove(4)
	if err != nil {
		t.Errorf("something went wrong")
	}

	err = l.Remove(7)
	if err != linkedlist.ErrPosOutOfRange {
		t.Errorf("something went wrong")
	}

	err = l.Remove(1)
	if err != nil {
		t.Errorf("something went wrong")
	}

	want := []int{1, 12}
	got := []int{}
	cur := l.Head
	got = append(got, cur.Data)

	for cur.Next != nil {
		cur = cur.Next
		if cur == l.Tail.Next {
			break
		}
		got = append(got, cur.Data)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v , want: %v", got, want)
	}
}
