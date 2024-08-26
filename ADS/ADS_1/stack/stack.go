package main

import (
	"fmt"
)

type Stack[T any] struct {
	stack []T
}

func (st *Stack[T]) Size() int {
	return len(st.stack)
}

func (st *Stack[T]) Peek() (T, error) {
	var result T
	if len(st.stack) == 0 {
		return result, fmt.Errorf("stack is empty")
	}
	result = st.stack[len(st.stack)-1]
	return result, nil
}

func (st *Stack[T]) Pop() (T, error) {
	var result T
	if len(st.stack) == 0 {
		return result, fmt.Errorf("stack is empty")
	}
	result = st.stack[len(st.stack)-1]
	st.stack = st.stack[:len(st.stack)-1]
	return result, nil
}

func (st *Stack[T]) Push(itm T) {
	st.stack = append(st.stack, itm)
}
func checkSt(str string) bool {
	var stack Stack[string]
	for _, chr := range str {
		s := string(chr)
		peek, _ := stack.Peek()
		if s == ")" && peek == "(" {
			stack.Pop()
		} else {
			stack.Push(s)
		}
	}
	return stack.Size() == 0
}
