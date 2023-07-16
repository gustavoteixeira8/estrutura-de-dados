package main

import (
	"errors"
	"fmt"
)

type Deque struct {
	firstPos int
	lastPos  int
	currLen  int
	maxLen   int
	vector   []int
}

func (d *Deque) dequeIsEmpty() bool {
	return d.firstPos == -1
}

func (d *Deque) dequeIsFull() bool {
	return (d.firstPos == 0 && d.lastPos == d.maxLen-1) || (d.firstPos == d.lastPos+1)
}

func (d *Deque) InsertLeft(v int) error {
	if d.dequeIsFull() {
		return errors.New("deque is full")
	}

	if d.firstPos == -1 {
		d.firstPos = 0
		d.lastPos = 0
	} else if d.firstPos == 0 {
		d.firstPos = d.maxLen - 1
	} else {
		d.firstPos -= 1
	}

	d.vector[d.firstPos] = v

	return nil
}

func (d *Deque) InsertRight(v int) error {
	if d.dequeIsFull() {
		return errors.New("deque is full")
	}

	if d.firstPos == -1 {
		d.firstPos = 0
		d.lastPos = 0
	} else if d.lastPos == d.maxLen-1 {
		d.lastPos = 0
	} else {
		d.lastPos += 1
	}

	d.vector[d.lastPos] = v

	return nil
}

func (d *Deque) RemoveLeft() error {
	if d.dequeIsEmpty() {
		return errors.New("deque is empty")
	}

	if d.firstPos == d.lastPos {
		d.firstPos = -1
		d.lastPos = -1
		return nil
	}

	if d.firstPos == d.maxLen-1 {
		d.firstPos = 0
		return nil
	}

	d.firstPos += 1

	return nil
}

func (d *Deque) RemoveRight() error {
	if d.dequeIsEmpty() {
		return errors.New("deque is empty")
	}

	if d.firstPos == d.lastPos {
		d.firstPos = -1
		d.lastPos = -1
		return nil
	}

	if d.firstPos == 0 {
		d.lastPos = d.maxLen - 1
		return nil
	}

	d.lastPos -= 1

	return nil
}

func (d *Deque) GetFirstPos() (int, error) {
	if d.dequeIsEmpty() {
		return -1, errors.New("deque is empty")
	}

	return d.vector[d.firstPos], nil
}

func (d *Deque) GetLastPos() (int, error) {
	if d.dequeIsEmpty() || d.lastPos < 0 {
		return -1, errors.New("deque is empty")
	}

	return d.vector[d.lastPos], nil
}

func New(maxLen int) *Deque {
	if maxLen <= 0 {
		maxLen = 10
	}
	return &Deque{currLen: -1, firstPos: -1, lastPos: -1, vector: make([]int, maxLen), maxLen: maxLen}
}

func main() {
	deque := New(10)

	deque.InsertRight(5)
	deque.InsertRight(8)
	deque.InsertLeft(3)
	deque.InsertLeft(4)

	deque.RemoveLeft()
	deque.RemoveRight()
	deque.RemoveRight()

	f, _ := deque.GetFirstPos()
	l, _ := deque.GetLastPos()

	fmt.Printf("First pos: %d\n", f)
	fmt.Printf("Last pos: %d\n", l)
	// fmt.Println(deque.vector)
}
