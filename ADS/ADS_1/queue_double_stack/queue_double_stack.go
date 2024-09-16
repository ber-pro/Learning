package main

import stack "ADS/ADS_1/stack"

type Queue[T any] struct {
	firstStack  stack.Stack[T]
	secondStack stack.Stack[T]
}

func (q *Queue[T]) Size() int {
	return q.firstStack.Size() + q.secondStack.Size()
}

func (q *Queue[T]) Dequeue() T {
	var result T
	if q.secondStack.Size() == 0 {
		for q.firstStack.Size() > 0 {
			val, _ := q.firstStack.Pop()
			q.secondStack.Push(val)
		}
	}
	if q.secondStack.Size() > 0 {
		result, _ = q.secondStack.Pop()
	}
	return result
}

func (q *Queue[T]) Enqueue(itm T) {
	q.firstStack.Push(itm)
}
