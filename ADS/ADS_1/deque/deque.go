package main

import "fmt"

type Deque[T any] struct {
	deque []T
}

func (d *Deque[T]) Size() int {
	return len(d.deque)
}

func (d *Deque[T]) AddFront(itm T) {
	d.deque = append(d.deque, itm)
	copy(d.deque[1:], d.deque)
	d.deque[0] = itm
}

func (d *Deque[T]) AddTail(itm T) {
	d.deque = append(d.deque, itm)
}

func (d *Deque[T]) RemoveFront() (T, error) {
	var result T
	if d.Size() == 0 {
		return result, fmt.Errorf("deque is empty")
	}
	result = d.deque[0]
	d.deque = d.deque[1:]
	return result, nil
}

func (d *Deque[T]) RemoveTail() (T, error) {
	var result T
	if d.Size() == 0 {
		return result, fmt.Errorf("deque is empty")
	}
	result = d.deque[d.Size()-1]
	d.deque = d.deque[:d.Size()-1]
	return result, nil
}
