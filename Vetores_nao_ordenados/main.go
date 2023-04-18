package main

import (
	"errors"
	"fmt"
)

type UnsortedVector struct {
	vector  []int
	lastPos int
	maxLen  int
}

func (v *UnsortedVector) Insert(val int) error {
	if v.lastPos == (v.maxLen - 1) {
		return errors.New("vector is full")
	}

	v.lastPos++
	v.vector[v.lastPos] = val
	return nil
}

func (v *UnsortedVector) Remove(val int) (int, error) {
	if v.lastPos == -1 {
		return -1, errors.New("vector is empty")
	}

	valIndex := v.Search(val)

	for i := valIndex; i > v.lastPos; i++ {
		v.vector[i] = v.vector[i+1]
	}

	v.lastPos--

	return valIndex, nil
}

func (v *UnsortedVector) Search(val int) int {
	index := -1

	for i, vectorVal := range v.vector {
		if vectorVal == val {
			index = i
			break
		}
	}

	return index
}

func (v *UnsortedVector) Vector() []int {
	return v.vector
}

func New(maxLen int) *UnsortedVector {
	if maxLen <= 0 {
		maxLen = 10
	}
	return &UnsortedVector{vector: make([]int, maxLen), lastPos: -1, maxLen: maxLen}
}

func main() {
	v := New(10)

	v.Insert(2) // 0
	v.Insert(4) // 1
	v.Insert(6) // 2
	v.Insert(8) // 3

	// fmt.Println(v.Search(8)) // 3

	v.Remove(8)

	v.Insert(10) // 3
	v.Insert(12) // 3
	v.Insert(14) // 3

	v.Insert(8)

	v.Insert(16) // 3
	v.Insert(18) // 3
	v.Insert(20) // 3

	v.Insert(22) // 3

	fmt.Println(v.Vector())
}
