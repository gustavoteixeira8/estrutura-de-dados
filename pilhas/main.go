package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	lastPos int
	vector  []int
	maxLen  int
}

func (s *Stack) Insert(val int) error {
	if s.lastPos == s.maxLen-1 {
		return errors.New("vector is full")
	}

	s.lastPos += 1
	s.vector[s.lastPos] = val

	return nil
}

func (s *Stack) Remove() error {
	if s.lastPos == -1 {
		return errors.New("vector is empty")
	}

	s.lastPos -= 1

	return nil
}

func (s *Stack) Tail() int {
	if s.lastPos == -1 {
		return -1
	}
	return s.vector[s.lastPos]
}

func (s *Stack) Vector() []int {
	return s.vector
}

func New(maxLen int) *Stack {
	return &Stack{vector: make([]int, maxLen), maxLen: maxLen, lastPos: -1}
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
	v.Remove()
	v.Insert(12)
	v.Remove()
	v.Insert(11)

	fmt.Println(v.Tail())
	fmt.Println(v.Vector())
}
