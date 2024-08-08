package main

import "fmt"

type DynamicArray[T any] struct {
	len      int
	capacity int
	array    []T
}

func (da *DynamicArray[T]) Init() {
	da.len = 0
	da.MakeDynamicArray(2)
}

func (da *DynamicArray[T]) MakeDynamicArray(size int) {
	var arr = make([]T, size)
	copy(arr, da.array)
	da.array = arr
	da.capacity = size
}

func (da *DynamicArray[T]) Insert(val T, index int) error {
	if index == da.len {
		da.Append(val)
		return nil
	}
	if index > da.len || index < 0 {
		return fmt.Errorf("Index out of range")
	}

	for i := da.len; i > index; i-- {
		da.array[i] = da.array[i-1]
	}
	da.len++
	da.array[index] = val
	return nil
}

func (da *DynamicArray[T]) Remove(index int) error {
	if da.len < index || index < 0 {
		return fmt.Errorf("Index out of range, index = %d", index)
	}

	for i := index; i < da.len; i++ {
		da.array[i] = da.array[i+1]
	}
	da.len--

	capacityPercent := float32(da.len) / float32(da.capacity)
	newCapacity := (da.capacity * 2) / 3

	if capacityPercent < 0.5 && newCapacity >= 8 {
		da.capacity = newCapacity
	}

	if capacityPercent < 0.5 && newCapacity < 8 {
		da.capacity = 8
	}
	return nil
}

func (da *DynamicArray[T]) Append(val T) {
	if da.len == da.capacity {
		da.MakeDynamicArray(da.capacity * 2)
	}
	da.array[da.len] = val
	da.len++
}

func (da *DynamicArray[T]) GetItem(index int) (T, error) {
	var result T
	if index > da.capacity {
		return result, fmt.Errorf("Index out of range, len array=%d", da.len)
	}
	result = da.array[index]
	return result, nil
}
