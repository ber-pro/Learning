package main

import (
	"reflect"
	"testing"
)

func TestCreateLinkedList(t *testing.T) {
	list := LinkedList{}

	list.AddInTail(Node{prev: nil, next: nil, val: 1})
	list.AddInTail(Node{prev: nil, next: nil, val: 2})
	list.AddInTail(Node{prev: nil, next: nil, val: 3})

	if list.head.val != 1 {
		t.Errorf("First value must be 1. However, first value is %d", list.head.val)
	}

	if list.head.next.val != 2 {
		t.Errorf("Second value must be 2. However, second value is %d", list.head.next.val)
	}

	if list.head.prev != nil {
		t.Errorf("The prev value of head must be nil")
	}

	if list.tail.val != 3 {
		t.Errorf("The tail value must be 3. However, the tail value is %d", list.tail.val)
	}

	if list.tail.prev.val != 2 {
		t.Errorf("The prev value of tail must be 2. However, the prev value of tail is %d", list.tail.prev.val)
	}
}

func TestInsertFirst(t *testing.T) {
	list := LinkedList{}

	list.InsertFirst(Node{prev: nil, next: nil, val: 3})
	list.InsertFirst(Node{prev: nil, next: nil, val: 2})
	list.InsertFirst(Node{prev: nil, next: nil, val: 1})

	if list.head.val != 1 {
		t.Errorf("First value must be 1. However, first value is %d", list.head.val)
	}

	if list.head.next.val != 2 {
		t.Errorf("Second value must be 2. However, second value is %d", list.head.next.val)
	}

	if list.head.prev != nil {
		t.Errorf("The prev value of head must be nil")
	}

	if list.tail.val != 3 {
		t.Errorf("The tail value must be 3. However, the tail value is %d", list.tail.val)
	}

	if list.tail.prev.val != 2 {
		t.Errorf("The prev value of tail must be 2. However, the prev value of tail is %d", list.tail.prev.val)
	}
}

func TestClean(t *testing.T) {
	list := LinkedList{}

	list.InsertFirst(Node{prev: nil, next: nil, val: 3})
	list.InsertFirst(Node{prev: nil, next: nil, val: 2})
	list.InsertFirst(Node{prev: nil, next: nil, val: 1})

	list.Clean()

	if list.head != nil || list.tail != nil {
		t.Error("Linked list don't clean")
	}
}

func TestInsertAfterHead(t *testing.T) {
	list := LinkedList{}

	node1 := Node{prev: nil, next: nil, val: 1}
	node2 := Node{prev: nil, next: nil, val: 2}
	node3 := Node{prev: nil, next: nil, val: 3}

	list.Insert(list.head, node1) // 1
	list.Insert(list.head, node2) // 1 2
	list.Insert(list.head, node3) // 1 3 2

	if list.head.val != 1 {
		t.Errorf("The head value must be 1 not %d", list.head.val)
	}

	if list.head.next.val != 3 {
		t.Errorf("The middle value must be 3 not %d", list.head.next.val)
	}

	if list.tail.val != 2 {
		t.Errorf("The tail value must be 2 not %d", list.tail.val)
	}
}

func TestFind(t *testing.T) {
	list := LinkedList{}

	node1 := Node{prev: nil, next: nil, val: 1}
	node2 := Node{prev: nil, next: nil, val: 2}
	node3 := Node{prev: nil, next: nil, val: 3}

	list.AddInTail(node1)
	list.AddInTail(node2)
	list.AddInTail(node3)

	firstFinsNode, err := list.Find(1)
	if err != nil {
		t.Errorf("Node with value = 1 %s", err)
	}

	if !reflect.DeepEqual(firstFinsNode, list.head) {
		t.Errorf("Found node not equal with head node of list")
	}

	if !reflect.DeepEqual(firstFinsNode.next, list.head.next) {
		t.Errorf("The next node of found node is not equal second node")
	}

	if !reflect.DeepEqual(firstFinsNode.next.next, list.tail) {
		t.Errorf("The found node is not equal list tail")
	}

}

func TestFindAll(t *testing.T) {
	list := LinkedList{}

	node1 := Node{prev: nil, next: nil, val: 1}
	node2 := Node{prev: nil, next: nil, val: 2}
	node3 := Node{prev: nil, next: nil, val: 3}
	node4 := Node{prev: nil, next: nil, val: 4}
	node5 := Node{prev: nil, next: nil, val: 5}

	list.AddInTail(node1)
	list.AddInTail(node2)
	list.AddInTail(node3)
	list.AddInTail(node1)
	list.AddInTail(node2)
	list.AddInTail(node3)
	list.AddInTail(node4)
	list.AddInTail(node5)

	findNodes := list.FindAll(1)
	if len(findNodes) == 0 {
		t.Errorf("List of found nodes is empty")
	}
	nodes := []Node{*list.head, *list.head.next.next.next}
	if !reflect.DeepEqual(findNodes, nodes) {
		t.Errorf("Found nodes not equal")
	}

	list2 := LinkedList{}
	list2.AddInTail(node1)
	list2.AddInTail(node1)
	list2.AddInTail(node1)
	list2.AddInTail(node1)

	findNodes = list2.FindAll(1)
	nodes = []Node{*list2.head, *list2.head.next, *list2.head.next.next, *list2.head.next.next.next}
	if !reflect.DeepEqual(findNodes, nodes) {
		t.Errorf("Found nodes not equal")
	}
}

func TestDelete(t *testing.T) {
	list := LinkedList{}

	node1 := Node{prev: nil, next: nil, val: 1}
	node2 := Node{prev: nil, next: nil, val: 2}
	node3 := Node{prev: nil, next: nil, val: 3}

	list.AddInTail(node1)
	list.AddInTail(node2)
	list.AddInTail(node3)
	list.AddInTail(node1)

	list.Delete(1, true)
	listEqual := LinkedList{}
	listEqual.AddInTail(node2)
	listEqual.AddInTail(node3)

	if !reflect.DeepEqual(list, listEqual) {
		t.Errorf("TEST 1: Lists not equal: %v %v", list.PrintListVals(), listEqual.PrintListVals())
	}

	if list.tail.val != listEqual.tail.val {
		t.Errorf("TEST 1: Last value must be %d not %d", listEqual.tail.val, list.tail.val)
	}

	if list.head.val != listEqual.head.val {
		t.Errorf("TEST 1: Head value must be %d not %d", list.head.val, listEqual.head.val)
	}
	if list.head.prev != listEqual.head.prev {
		t.Errorf("TEST 1: Prev first value must be %v not %v", list.head.prev, listEqual.head.prev)
	}

	node1 = Node{prev: nil, next: nil, val: 1}
	node2 = Node{prev: nil, next: nil, val: 1}
	node3 = Node{prev: nil, next: nil, val: 1}
	node4 := Node{prev: nil, next: nil, val: 1}
	node5 := Node{prev: nil, next: nil, val: 1}

	list2 := LinkedList{}
	list2.AddInTail(node1)
	list2.AddInTail(node2)
	list2.AddInTail(node3)
	list2.AddInTail(node4)
	list2.AddInTail(node5)

	list2.Delete(1, false)

	listEqual = LinkedList{}
	listEqual.AddInTail(node1)
	listEqual.AddInTail(node2)
	listEqual.AddInTail(node3)
	listEqual.AddInTail(node4)

	if !reflect.DeepEqual(list2, listEqual) {
		t.Errorf("TEST 2: List must be %v, is not %v", listEqual.PrintListVals(), list2.PrintListVals())
	}

	list2.Delete(1, true)
	listEqual = LinkedList{}
	if !reflect.DeepEqual(list2, listEqual) {
		t.Errorf("TEST 2: List must be %v, is not %v", listEqual.PrintListVals(), list2.PrintListVals())
	}

}
