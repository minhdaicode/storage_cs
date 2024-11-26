package stack_test

import (
	"reflect"
	"storage_cs/datastruture/stack"
	"testing"
)

func TestSliceStack(t *testing.T) {
	s := stack.NewSliceStack[int]()
	if !s.IsEmpty() {
		t.Errorf("something went wrong")
	}

	s.Push(5)
	s.Push(4)
	s.Push(3)
	s.Push(2)

	want := 2
	got, err := s.Peek()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v , want: %v", got, want)
	}
	if err != nil {
		t.Errorf("something went wrong")
	}

	want = 4
	got = int(s.Size())
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v , want: %v", got, want)
	}

	want = 2
	got, err = s.Pop()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v , want: %v", got, want)
	}
	if err != nil {
		t.Errorf("something went wrong")
	}

	want = 3
	got, err = s.Pop()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v , want: %v", got, want)
	}
	if err != nil {
		t.Errorf("something went wrong")
	}

	s.Clear()
	want = 0
	got, err = s.Pop()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v , want: %v", got, want)
	}
	if err != stack.ErrIsEmpty {
		t.Errorf("something went wrong")
	}

	got, err = s.Peek()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v , want: %v", got, want)
	}
	if err != stack.ErrIsEmpty {
		t.Errorf("something went wrong")
	}
}
