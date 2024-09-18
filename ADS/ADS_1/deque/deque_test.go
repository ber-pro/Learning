package main

import (
	"testing"
)

func TestDeque_AddFront(t *testing.T) {
	var deque Deque[int]

	deque.AddFront(1)
	deque.AddFront(2)
	deque.AddFront(3)

	if deque.Size() != 3 {
		t.Errorf("size of deque should be is 3? not is %d", deque.Size())
	}

	if value, _ := deque.RemoveFront(); value != 3 {
		t.Errorf("Value should be is 3, not is %d", value)
	}

	if value, _ := deque.RemoveTail(); value != 1 {
		t.Errorf("value should be 1, not is %d", value)
	}

	if value, _ := deque.RemoveFront(); value != 2 {
		t.Errorf("Value should be is 2, not is %d", value)
	}

	if _, err := deque.RemoveFront(); err == nil {
		t.Errorf("Should be error, deque is empty")
	}
}

func TestDeque_AddTail(t *testing.T) {
	var deque Deque[int]

	deque.AddTail(1)
	deque.AddTail(2)
	deque.AddTail(3)
	deque.AddTail(4)

	if value, _ := deque.RemoveFront(); value != 1 {
		t.Errorf("value should be 1, not is %d", value)
	}

	if value, _ := deque.RemoveTail(); value != 4 {
		t.Errorf("value should be 4, not is %d", value)
	}
}
