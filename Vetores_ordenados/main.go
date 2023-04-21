package main

import (
	"errors"
	"fmt"
	"math"
)

type SortedVector struct {
	vector  []int
	lastPos int
	maxLen  int
}

// O(n)
func (v *SortedVector) Insert(val int) error {
	if v.lastPos > v.maxLen-1 {
		return errors.New("vector is full")
	}

	pos := 0
	for i := 0; i <= v.lastPos; i++ {
		pos = i
		if val < v.vector[i] {
			break
		}
		if i == v.lastPos {
			pos = i + 1
		}
	}

	for i := v.lastPos; i >= pos; i-- {
		// pos = 3 | i = 3
		//  0  1  2  3  4  5
		// [0, 1, 2, 3, 4, 5, x, x, x]
		v.vector[i+1] = v.vector[i]
	}

	v.vector[pos] = val
	v.lastPos += 1

	return nil
}

func (v *SortedVector) Remove(val int) error {
	if v.lastPos == -1 {
		return errors.New("vector is empty")
	}

	//  0  1  2  3  4  5
	// [0, 1, 2, 3, 4, 5, x, x, x]

	pos := v.LinearSearch(val)

	for i := pos; i < v.lastPos; i++ {
		v.vector[i] = v.vector[i+1]
	}

	v.lastPos -= 1

	return nil
}

// O(n)
func (v *SortedVector) LinearSearch(val int) int {
	pos := -1
	for i, vectorValue := range v.vector {
		if vectorValue > val {
			break
		}
		if vectorValue == val {
			pos = i
			break
		}
	}

	return pos
}

func (v *SortedVector) BinarySearch(val int) int {
	head := 0
	tail := v.lastPos

	//  0  1  2  3  4  5  6
	// [1  2  4  5  9  11 12 0 0 0]

	for head <= tail {
		middle := int(math.Floor(float64(head+tail) / 2))
		if v.vector[middle] == val {
			return middle
		}
		if v.vector[middle] > val {
			tail = middle - 1
		} else {
			head = middle + 1
		}
	}

	return -1
}

func (v *SortedVector) Vector() []int {
	return v.vector
}

func New(maxLen int) *SortedVector {
	return &SortedVector{vector: make([]int, maxLen), maxLen: maxLen, lastPos: -1}
}

func main() {
	v := New(10)

	v.Insert(10)
	v.Insert(9)
	v.Insert(1)
	v.Insert(2)
	v.Insert(4)
	v.Insert(3)
	v.Insert(5)
	v.Remove(3)
	v.Insert(12)
	v.Remove(10)
	v.Insert(11)

	fmt.Println(v.BinarySearch(12))

	fmt.Println(v.Vector())

}
