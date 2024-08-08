package main

import (
	"testing"
)

func TestDynamicArray_MakeDynamicArray(t *testing.T) {
	var testArray DynamicArray[int]
	const InitialCapacity = 2
	const InitialLen = 0

	testArray.Init()

	if testArray.capacity != InitialCapacity {
		t.Errorf("Initial array capacite must be %d, not is %d", InitialCapacity, testArray.capacity)
	}
	if testArray.len != InitialLen {
		t.Errorf("Initial array len must be %d, not is %d", InitialLen, testArray.len)
	}

	testArray.Append(1)
	testArray.Append(2)

	if testArray.capacity != InitialCapacity {
		t.Errorf("Initial array capacite must be %d, not is %d", InitialCapacity, testArray.capacity)
	}

	testArray.Append(3)

	if testArray.capacity != InitialCapacity*2 {
		t.Errorf("Initial array capacite must be %d, not is %d", 4, testArray.capacity)
	}

	if testArray.len != 3 {
		t.Errorf("Initial array len must be %d, not is %d", 3, testArray.len)
	}

	if val, _ := testArray.GetItem(2); val != 3 {
		t.Errorf("Value should be %d not is %d", 3, val)
	}
	if val, _ := testArray.GetItem(1); val != 2 {
		t.Errorf("Value should be %d not is %d", 2, val)
	}
	if val, _ := testArray.GetItem(0); val != 1 {
		t.Errorf("Value should be %d not is %d", 1, val)
	}
}

func TestDynamicArray_Insert(t *testing.T) {
	var testArray DynamicArray[int]
	testArray.Init()
	_ = testArray.Insert(1, 0)

	if testArray.array[0] != 1 {
		t.Errorf("Value of array in index")
	}

	_ = testArray.Insert(2, 0)
	if testArray.array[0] != 2 {
		t.Errorf("Wrong value of %d indexed element, should be %d", 0, 2)
	}
	if testArray.array[1] != 1 {
		t.Errorf("Wrong value of %d indexed element, should be %d", 0, 1)
	}
}

func TestDynamicArray_RelocateAndGetItem(t *testing.T) {
	var testArray DynamicArray[int]
	const NumberIter = 16
	testArray.Init()
	for i := 1; i <= NumberIter; i++ {
		testArray.Append(i)
	}

	testArray.Append(88)
	if testArray.capacity != 32 {
		t.Errorf("Initial array capacite must be %d, not is %d", 32, testArray.capacity)
	}
	if testArray.len != 17 {
		t.Errorf("Initial array len must be %d, not is %d", 17, testArray.len)
	}

	for i := 1; i <= 13; i++ {
		err := testArray.Remove(0)
		if err != nil {
			t.Errorf("%s", err)
		}
	}

	if testArray.capacity != 8 {
		t.Errorf("Initial array capacite must be %d, not is %d", 8, testArray.capacity)
	}
	if testArray.len != 4 {
		t.Errorf("Initial array len must be %d, not is %d", 4, testArray.len)
	}

	getValue, err := testArray.GetItem(3)
	if err != nil {
		t.Errorf("Error while use func GetItem with %d index", 3)
	}
	if getValue != 88 {
		t.Errorf("Wrong of implementation GetItem getValue=%d should be %d", getValue, 88)
	}
	getValue, err = testArray.GetItem(0)
	if err != nil {
		t.Errorf("Error while use func GetItem with %d index", 0)
	}
	if getValue != 14 {
		t.Errorf("Wrong of implementation GetItem")
	}

}
