package main

import (
	"constraints"
	"fmt"
)

type Node[T constraints.Ordered] struct {
	prev  *Node[T]
	next  *Node[T]
	value T
}

type OrderedList[T constraints.Ordered] struct {
	head       *Node[T]
	tail       *Node[T]
	_ascending bool
}

func (l *OrderedList[T]) Count() int {
	counterNode := 0
	currentNode := l.head
	for currentNode != nil {
		counterNode++
		currentNode = currentNode.next
	}
	return counterNode
}

func (l *OrderedList[T]) Add(item T) {
	if l.head == nil {
		var newNode Node[T]
		newNode.value = item
		l.head = &newNode
		l.tail = &newNode
	} else {
		var newNode Node[T]
		newNode.value = item
		currentNode := l.head
		for currentNode != nil {
			if (l._ascending == true && l.Compare(newNode.value, currentNode.value) <= 0) || (l._ascending == false && l.Compare(newNode.value, currentNode.value) >= 0) {
				//случай когда нашли нужное место
				// 1 - это начало списка
				if currentNode.prev == nil {
					newNode.next = currentNode
					currentNode.prev = &newNode
					l.head = &newNode
					// 2 - это середина списка
				} else if currentNode.prev != nil {
					currentNode.prev.next = &newNode
					newNode.prev = currentNode.prev
					newNode.next = currentNode
					currentNode.prev = &newNode
				}
				break
			} else if currentNode.next == nil {
				// 3 - это конец списка
				currentNode.next = &newNode
				newNode.prev = currentNode
				l.tail = &newNode
				break
			}
			currentNode = currentNode.next
		}
	}
}

func (l *OrderedList[T]) Find(n T) (Node[T], error) {
	var err error
	if l.head == nil {
		err = fmt.Errorf("list is empty")
		return Node[T]{value: n, next: nil, prev: nil}, err
	}

	currentNode := l.head
	for currentNode != nil {
		if (l._ascending == true && l.Compare(n, currentNode.value) < 0) || (l._ascending == false && l.Compare(n, currentNode.value) > 0) {
			err = fmt.Errorf("value = %T not found in list", n)
			return Node[T]{value: n, next: nil, prev: nil}, err
		} else if currentNode.value == n {
			return *currentNode, nil
		}
		currentNode = currentNode.next
	}
	err = fmt.Errorf("value not found")
	return Node[T]{value: n, next: nil, prev: nil}, err
}

func (l *OrderedList[T]) Delete(n T) {
	currentNode := l.head
	for currentNode != nil {
		if currentNode.value == n && currentNode.prev == nil {
			if l.head == l.tail {
				l.Clear(l._ascending)
				break
			}
			l.head = l.head.next
			l.head.prev = nil
			break
		} else if currentNode.value == n && currentNode.next == nil {
			l.tail = l.tail.prev
			l.tail.next = nil
			break
		} else if currentNode.value == n {
			currentNode.prev.next = currentNode.next
			currentNode.next.prev = currentNode.prev
			break
		}
		currentNode = currentNode.next
	}
}

func (l *OrderedList[T]) Clear(asc bool) {
	l.head = nil
	l.tail = nil
	l._ascending = asc
}

func (l *OrderedList[T]) Compare(v1 T, v2 T) int {
	if v1 < v2 {
		return -1
	}
	if v1 > v2 {
		return +1
	}
	return 0
}
