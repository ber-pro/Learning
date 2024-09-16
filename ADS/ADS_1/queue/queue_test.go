package main

import "testing"

func TestMethodQueue(t *testing.T) {
	var queue Queue[int]

	size := queue.Size()
	if size != 0 {
		t.Errorf("Size queue should be 0, not is %d", size)
	}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	size = queue.Size()
	if size != 3 {
		t.Errorf("Size queue should be 3, not is %d", size)
	}

	val, err := queue.Dequeue()
	if err != nil {
		t.Errorf("error in implementation")
	}
	if val != 1 {
		t.Errorf("value should be 1")
	}
}
