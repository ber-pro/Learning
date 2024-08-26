package main

import (
	"testing"
)

func TestMethod(t *testing.T) {
	var stack Stack[int]

	stackSize := stack.Size()
	if stackSize != 0 {
		t.Errorf("Init size of stack should be 0, not is %d", stackSize)
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	stackSize = stack.Size()
	if stackSize != 3 {
		t.Errorf("Init size of stack should be 3, not is %d", stackSize)
	}

	itm, err := stack.Pop()
	if err != nil {
		t.Errorf("Current item of stack should be 3, not is %s", err)
	}
	if itm != 3 {
		t.Errorf("Current item of stack should be 3, not is %d", itm)
	}

	itm, err = stack.Pop()
	if err != nil {
		t.Errorf("Current item of stack should be 2, not is %s", err)
	}
	if itm != 2 {
		t.Errorf("Current item of stack should be 2, not is %d", itm)
	}

	itm, err = stack.Pop()
	if err != nil {
		t.Errorf("Current item of stack should be 1, not is %s", err)
	}
	if itm != 1 {
		t.Errorf("Current item of stack should be 1, not is %d", itm)
	}

	itm, err = stack.Pop()
	if err == nil {
		t.Errorf("stack should be empty")
	}
}

func TestCheckSt(t *testing.T) {
	string_1 := "(())()(())"   // True
	string_2 := "(()((())()))" // True
	string_3 := "))(("         // False
	string_4 := "((())"        // False
	string_5 := "(()))"        // False

	res := checkSt(string_1)
	if res != true {
		t.Errorf("Error in test 1: %s", string_1)
	}

	res = checkSt(string_2)
	if res != true {
		t.Errorf("Error in test 2: %s", string_2)
	}

	res = checkSt(string_3)
	if res != false {
		t.Errorf("Error in test 3: %s", string_3)
	}

	res = checkSt(string_4)
	if res != false {
		t.Errorf("Error in test 4: %s", string_4)
	}

	res = checkSt(string_5)
	if res != false {
		t.Errorf("Error in test 5: %s", string_5)
	}

}
