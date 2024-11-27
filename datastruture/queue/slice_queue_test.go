package queue_test

import (
	"fmt"
	"reflect"
	"storage_cs/datastruture/queue"
	"testing"
)

func TestSliceQueue(t *testing.T) {
	q := queue.NewSliceQueue[int]()
	if !q.IsEmpty() {
		t.Errorf("something went wrong")
	}

	want := 0
	got, err := q.Front()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v , want: %v", got, want)
	}
	if err != queue.ErrIsEmpty {
		t.Errorf("something went wrong")
	}

	got, err = q.Rear()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v , want: %v", got, want)
	}
	if err != queue.ErrIsEmpty {
		t.Errorf("something went wrong")
	}

	got, err = q.DeQueue()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v , want: %v", got, want)
	}
	if err != queue.ErrIsEmpty {
		t.Errorf("something went wrong")
	}

	q.EnQueue(5)
	q.EnQueue(4)
	q.EnQueue(3)
	q.EnQueue(2)

	want = 5
	got, err = q.Front()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v , want: %v", got, want)
	}
	if err != nil {
		t.Errorf("something went wrong")
	}

	want = 4
	got = int(q.Size())
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v , want: %v", got, want)
	}

	want = 5
	got, err = q.DeQueue()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v , want: %v", got, want)
	}
	if err != nil {
		t.Errorf("something went wrong")
	}

	want = 2
	got, err = q.Rear()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v , want: %v", got, want)
	}
	if err != nil {
		t.Errorf("something went wrong")
	}

	fmt.Println(q.String())
}
