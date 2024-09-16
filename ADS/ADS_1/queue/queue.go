package main

import "fmt"

type Queue[T any] struct {
	queue []T
}

func (q *Queue[T]) Size() int {
	var sizeQueue int
	sizeQueue = len(q.queue)
	return sizeQueue
}

func (q *Queue[T]) Dequeue() (T, error) {
	var result T
	if q.Size() > 0 {
		result = q.queue[0]
		q.queue = q.queue[1:]
	} else {
		return result, fmt.Errorf("queue is empty")
	}
	return result, nil
}

func (q *Queue[T]) Enqueue(itm T) {
	q.queue = append(q.queue, itm)
}
