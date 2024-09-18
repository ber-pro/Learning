package main

import "testing"

func TestOrderedListAsc(t *testing.T) {
	var l OrderedList[int]
	l._ascending = true

	l.Add(1)

	if l.head.value != 1 || l.tail.value != 1 {
		t.Errorf("wrong add node, list should be [1]")
	}

	l.Add(2)
	if l.head.value != 1 || l.tail.value != 2 {
		t.Errorf("wrong add node, list should be [1 2]")
	}

	if l.Count() != 2 {
		t.Errorf("wrong count node, should be 2")
	}

	l.Add(10)
	if l.head.value != 1 || l.head.next.value != 2 || l.tail.prev.value != 2 || l.tail.value != 10 {
		t.Errorf("wrong add node, list should be [1 2 10]")
	}

	_, err := l.Find(10)
	if err != nil {
		t.Errorf("wrong find, [1 2 10]")
	}

	_, err = l.Find(1)
	if err != nil {
		t.Errorf("wrong find, [1 2 10]")
	}
	_, err = l.Find(2)
	if err != nil {
		t.Errorf("wrong find, [1 2 10]")
	}

	l.Delete(2)
	if l.head.next.value != 10 || l.tail.prev.value != 1 {
		t.Errorf("wrong delete, [1 10]")
	}

	l.Delete(10)
	if l.head.value != l.tail.value {
		t.Errorf("wrong delete, [1]")
	}

	l.Delete(1)
	if l.head != nil || l.tail != nil {
		t.Errorf("wrong delete, [ ]")
	}

}

func TestOrderedListDesc(t *testing.T) {
	var l OrderedList[int]
	l._ascending = false

	l.Add(1)
	l.Add(5)
	l.Add(10)
	l.Add(15)

	if l.Count() != 4 {
		t.Errorf("wrong count node")
	}

	_, err := l.Find(20)
	if err == nil {
		t.Errorf("Attemd to find something not equaul [1 5 10 15]")
	}

	node, err := l.Find(1)
	if err != nil && node.value != 1 {
		t.Errorf("Attemd to find something equaul [1 5 10 15]")
	}

	l.Delete(1)
	if l.tail.value != 5 || l.tail.next != nil {
		t.Errorf("wrong delete tail in list")
	}

	l.Delete(15)
	if l.head.value != 10 || l.head.prev != nil {
		t.Errorf("wrong delete head in list")
	}

}
