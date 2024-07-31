package main

import (
	"fmt"
	"reflect"
)

type Node struct {
	prev *Node
	next *Node
	val  int
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) AddInTail(node Node) {
	if l.head == nil {
		l.head = &node
		l.head.prev = nil
		l.head.next = nil
	} else {
		l.tail.next = &node
		node.prev = l.tail
	}
	l.tail = &node
	l.tail.next = nil
}

func (l *LinkedList) Insert(after *Node, node Node) {
	if reflect.DeepEqual(l.tail, after) {
		l.AddInTail(node)
	} else {
		node.prev = after
		node.next = after.next
		after.next = &node
	}
}

func (l *LinkedList) InsertFirst(node Node) {
	if l.head == nil {
		l.AddInTail(node)
	} else {
		node.next = l.head
		l.head.prev = &node
		l.head = &node
	}
}

func (l *LinkedList) PrintListVals() []int {
	currentNode := l.head
	listValues := []int{}
	for currentNode != nil {
		listValues = append(listValues, currentNode.val)
		currentNode = currentNode.next
	}

	return listValues
}

func (l *LinkedList) Find(value int) (*Node, error) {
	currentNode := l.head
	for currentNode != nil {
		if currentNode.val == value {
			return currentNode, nil
		}
		currentNode = currentNode.next
	}
	return &Node{prev: nil, next: nil, val: -1}, fmt.Errorf("not found")
}

func (l *LinkedList) FindAll(value int) []Node {
	findNodes := []Node{}
	currentNode := l.head
	for currentNode != nil {
		if currentNode.val == value {
			findNodes = append(findNodes, *currentNode)
		}
		currentNode = currentNode.next
	}
	return findNodes
}

func (l *LinkedList) Delete(val int, isAll bool) {
	currentNode := l.head
	if isAll {
		for currentNode != nil {
			if currentNode.val == val {
				if reflect.DeepEqual(currentNode, l.head) {
					if reflect.DeepEqual(currentNode, l.tail) {
						l.Clean()
						return
					}
					l.head = l.head.next
					l.head.prev = nil
				} else if reflect.DeepEqual(currentNode, l.tail) {
					l.tail = currentNode.prev
					l.tail.next = nil
				} else {
					currentNode.prev.next = currentNode.next
				}
			}
			currentNode = currentNode.next
		}
		return
	}
	for currentNode != nil {
		if currentNode.val == val {
			if reflect.DeepEqual(currentNode, l.head) {
				if reflect.DeepEqual(currentNode, l.tail) {
					l.Clean()
					break
				}
				l.head = l.head.next
				l.head.prev = nil
				break
			} else {
				currentNode.prev.next = currentNode.next
			}
		}
		currentNode = currentNode.next
	}

}

func (l *LinkedList) Clean() {
	l.head = nil
	l.tail = nil
}
